import {
  Box,
  IconButton,
  Container,
  Divider,
  Grid,
  Typography,
  useMediaQuery,
  TextField,
  Autocomplete,
  LinearProgress,
} from "@mui/material";
import { useContext, useEffect, useState } from "react";
import {
  BlueprintInput,
  ListGroupsQuery,
  ProvidersQuery,
  useCreateBlueprintMutation,
  useGetBlueprintLazyQuery,
  useListGroupsQuery,
  useProvidersQuery,
  useUpdateBlueprintMutation,
} from "../../api/graphql/generated";
import MonacoEditor, { useMonaco } from "@monaco-editor/react";
import { configureMonacoYaml } from "monaco-yaml";
import { ThemeContext } from "../../theme";
import { Add, AutoFixHigh, Circle, Save } from "@mui/icons-material";
import { MuiMarkdown } from "mui-markdown";
import { LoadingButton } from "@mui/lab";
import { useSnackbar } from "notistack";
import { useNavigate, useParams } from "react-router-dom";

export default function BlueprintForm({ action }: { action: "create" | "edit" }) {
  const { id } = useParams();
  const { themePreference } = useContext(ThemeContext);
  const prefersDarkMode = useMediaQuery("(prefers-color-scheme: dark)");
  const [blueprint, setBlueprint] = useState<BlueprintInput>({
    name: "",
    description: "### Untitled Blueprint",
    blueprintTemplate: `version: "1.0"\n`,
    providerId: "",
    parentGroupId: "",
  });
  const monaco = useMonaco();
  const { data: providersData, error: providersError, loading: providersLoading } = useProvidersQuery();
  const { data: groupsData, error: groupsError, loading: groupsLoading } = useListGroupsQuery();
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  // Create
  const [createBlueprint, { data: createBlueprintData, error: createBlueprintError, loading: createBlueprintLoading }] =
    useCreateBlueprintMutation();
  // Edit
  const [getBlueprint, { data: getBlueprintData, error: getBlueprintError, loading: getBlueprintLoading }] = useGetBlueprintLazyQuery();
  const [
    updateBlueprint,
    { data: updateBlueprintData, error: updateBlueprintError, loading: updateBlueprintLoading, reset: resetUpdateBlueprint },
  ] = useUpdateBlueprintMutation();

  useEffect(() => {
    if (providersError) enqueueSnackbar({ message: `Failed to load provider list: ${providersError.message}`, variant: "error" });
    if (groupsError) enqueueSnackbar({ message: `Failed to load group list: ${groupsError.message}`, variant: "error" });
    if (createBlueprintError) enqueueSnackbar({ message: `Failed to create blueprint: ${createBlueprintError.message}`, variant: "error" });
    if (getBlueprintError) enqueueSnackbar({ message: `Failed to create blueprint: ${getBlueprintError.message}`, variant: "error" });
    if (updateBlueprintError) enqueueSnackbar({ message: `Failed to update blueprint: ${updateBlueprintError.message}`, variant: "error" });
  }, [providersError, groupsError, createBlueprintError, getBlueprintError, updateBlueprintError]);

  // If we're editing the blueprint, fetch based on ID
  useEffect(() => {
    if (action === "edit" && id)
      getBlueprint({
        variables: { id },
      }).catch(console.error);
  }, [id, action, getBlueprint]);

  // Once we get edit data, update local form state
  useEffect(() => {
    if (getBlueprintData?.blueprint)
      setBlueprint({
        name: getBlueprintData.blueprint.name,
        description: getBlueprintData.blueprint.description,
        blueprintTemplate: getBlueprintData.blueprint.blueprintTemplate,
        parentGroupId: getBlueprintData.blueprint.parentGroup.id,
        providerId: getBlueprintData.blueprint.provider.id,
      } as BlueprintInput);
  }, [getBlueprintData]);

  // Move to edit page after creation
  useEffect(() => {
    if (createBlueprintData?.createBlueprint.id) {
      enqueueSnackbar({ message: "Create blueprint!", variant: "success" });
      navigate(`/blueprints/edit/${createBlueprintData.createBlueprint.id}`);
    }
  }, [createBlueprintData, navigate]);

  // Show pop up when updating blueprint
  useEffect(() => {
    if (updateBlueprintData?.updateBlueprint.id) {
      enqueueSnackbar({ message: "Updated blueprint!", variant: "success" });
      resetUpdateBlueprint();
    }
  }, [updateBlueprintData, enqueueSnackbar, resetUpdateBlueprint]);

  const handlePrettify = () => {
    if (monaco)
      monaco.editor.getEditors().forEach((editor) => {
        editor.getAction("editor.action.formatDocument")?.run().catch(console.error);
      });
  };

  const handleSubmitBlueprint = () => {
    if (action === "create")
      createBlueprint({
        variables: {
          input: blueprint,
        },
      }).catch(console.error);
    else if (action === "edit" && id)
      updateBlueprint({
        variables: {
          id,
          input: blueprint,
        },
      }).catch(console.error);
  };

  return (
    <Container sx={{ display: "flex", flexDirection: "column", height: "100%", py: 2 }}>
      <Box sx={{ display: "flex", alignItems: "center", justifyContent: "space-between" }}>
        <TextField
          variant="standard"
          placeholder="Untitled Blueprint"
          value={blueprint.name}
          onChange={(e) => setBlueprint((prev) => ({ ...prev, name: e.target.value }))}
        ></TextField>
        <Autocomplete
          disablePortal
          autoComplete
          clearOnEscape
          options={groupsData?.groups ?? []}
          getOptionKey={(option: ListGroupsQuery["groups"][0]) => `${option.id}`}
          getOptionLabel={(option: ListGroupsQuery["groups"][0]) => `${option.name}`}
          sx={{ width: 300 }}
          value={(groupsData?.groups.find((g) => g.id === blueprint.parentGroupId) || null) ?? null}
          onChange={(_, val) => {
            val && setBlueprint((prev) => ({ ...prev, parentGroupId: val.id }));
          }}
          renderInput={(params) => <TextField {...params} label="Parent Group" />}
          disabled={groupsLoading || groupsError !== undefined}
        />
        <Autocomplete
          disablePortal
          autoComplete
          clearOnEscape
          options={providersData?.providers ?? []}
          getOptionKey={(option: ProvidersQuery["providers"][0]) => `${option.id}`}
          getOptionLabel={(option: ProvidersQuery["providers"][0]) => `${option.displayName} (${option.providerVersion})`}
          sx={{ width: 300 }}
          value={(providersData?.providers.find((p) => p.id === blueprint.providerId) || null) ?? null}
          onChange={(_, val) => {
            val && setBlueprint((prev) => ({ ...prev, providerId: val.id }));
          }}
          renderOption={(props, option) => (
            <li {...props}>
              <Circle color={option.isLoaded ? "success" : "error"} sx={{ height: "0.75rem" }} />
              {option.displayName} ({option.providerVersion})
            </li>
          )}
          renderInput={(params) => <TextField {...params} label="Provider" />}
          disabled={providersLoading || providersError !== undefined}
        />
        <Box>
          <IconButton onClick={handlePrettify} sx={{ mr: 1 }} title="Prettify">
            <AutoFixHigh />
          </IconButton>
          <LoadingButton
            variant="contained"
            startIcon={action === "create" ? <Add /> : <Save />}
            loading={createBlueprintLoading || updateBlueprintLoading}
            disabled={
              blueprint.name === "" ||
              blueprint.blueprintTemplate === "" ||
              blueprint.parentGroupId === "" ||
              blueprint.providerId === "" ||
              createBlueprintData != null
            }
            onClick={handleSubmitBlueprint}
          >
            {action === "create" ? "Create" : "Save"}
          </LoadingButton>
        </Box>
      </Box>
      <Divider sx={{ my: 2 }} />
      {getBlueprintLoading && <LinearProgress />}
      <Grid container sx={{ flexGrow: 1 }} spacing={2}>
        <Grid item md={6}>
          <Typography variant="h6">Description</Typography>
          <Box sx={{ maxHeight: "35dvh", height: "50%", borderWidth: 1, borderStyle: "solid", borderColor: "primary" }}>
            <MonacoEditor
              theme={
                themePreference === "auto" ? (prefersDarkMode ? "vs-dark" : "light") : themePreference === "dark" ? "vs-dark" : "light"
              }
              language="markdown"
              options={{ minimap: { enabled: false }, scrollBeyondLastLine: false, wordWrap: "on" }}
              value={blueprint.description}
              onChange={(value, _) => setBlueprint((prev) => ({ ...prev, description: value ?? "" }))}
            ></MonacoEditor>
          </Box>
          <Divider sx={{ width: "100%", my: 2 }} />
          <Typography variant="h6">Preview:</Typography>
          <Box sx={{ maxHeight: "35dvh", overflowY: "scroll" }}>
            <MuiMarkdown>{blueprint.description}</MuiMarkdown>
          </Box>
        </Grid>
        <Grid item md={6}>
          <Box sx={{ maxHeight: "calc(100dvh - 12rem)", height: "100%", borderWidth: 1, borderStyle: "solid", borderColor: "primary" }}>
            <MonacoEditor
              theme={
                themePreference === "auto" ? (prefersDarkMode ? "vs-dark" : "light") : themePreference === "dark" ? "vs-dark" : "light"
              }
              language="yaml"
              options={{ minimap: { enabled: false }, scrollBeyondLastLine: false, wordWrap: "on" }}
              value={blueprint.blueprintTemplate}
              onChange={(value, _) => setBlueprint((prev) => ({ ...prev, blueprintTemplate: value ?? "" } as BlueprintInput))}
              beforeMount={(monaco) => {
                // @ts-ignore-error Monaco isn't of exact expected type, but still works
                configureMonacoYaml(monaco, {
                  enableSchemaRequest: true,
                  schemas: [
                    {
                      fileMatch: ["*"],
                      schema: {
                        type: "object",
                        properties: {
                          version: {
                            type: "string",
                            description: "The blueprint syntax version",
                          },
                        },
                      },
                      uri: "https://github.com/cble-platform",
                    },
                  ],
                });
              }}
            />
          </Box>
        </Grid>
      </Grid>
    </Container>
  );
}
