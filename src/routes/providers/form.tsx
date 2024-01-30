import { AutoFixHigh, Add, Save, Power, PowerOff, Cached, Circle } from "@mui/icons-material";
import { LoadingButton } from "@mui/lab";
import {
  Autocomplete,
  Box,
  ButtonGroup,
  Container,
  Divider,
  Grid,
  IconButton,
  LinearProgress,
  TextField,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { useNavigate, useParams } from "react-router-dom";
import {
  ProviderInput,
  useConfigrueProviderMutation,
  useCreateProviderMutation,
  useGetProviderLazyQuery,
  useLoadProviderMutation,
  useUnloadProviderMutation,
  useUpdateProviderMutation,
} from "../../api/generated";
import { useContext, useEffect, useState } from "react";
import { ThemeContext } from "../../theme";
import MonacoEditor, { useMonaco } from "@monaco-editor/react";
import { enqueueSnackbar } from "notistack";

type GitTag = {
  ref: string;
  node_id: string;
  url: string;
  object: {
    sha: string;
    type: string;
    url: string;
  };
};

export default function ProviderForm({ action }: { action: "edit" | "create" }) {
  const { id } = useParams();
  const navigate = useNavigate();
  const { themePreference } = useContext(ThemeContext);
  const prefersDarkMode = useMediaQuery("(prefers-color-scheme: dark)");
  const monaco = useMonaco();
  const [provider, setProvider] = useState<ProviderInput>({
    configBytes: "",
    displayName: "",
    providerGitUrl: "",
    providerVersion: "",
  });
  const [gitTags, setGitTags] = useState<GitTag[]>([]);
  // Create
  const [createProvider, { data: createProviderData, error: createProviderError, loading: createProviderLoading }] =
    useCreateProviderMutation();
  // Edit
  const [getProvider, { data: getProviderData, error: getProviderError, loading: getProviderLoading }] = useGetProviderLazyQuery();
  const [
    updateProvider,
    { data: updateProviderData, error: updateProviderError, loading: updateProviderLoading, reset: resetUpdateProvider },
  ] = useUpdateProviderMutation();
  const [
    configureProvider,
    { data: configureProviderData, error: configureProviderError, loading: configureProviderLoading, reset: resetConfigureProvider },
  ] = useConfigrueProviderMutation();
  const [loadProvider, { data: loadProviderData, error: loadProviderError, loading: loadProviderLoading, reset: resetLoadProvider }] =
    useLoadProviderMutation();
  const [
    unloadProvider,
    { data: unloadProviderData, error: unloadProviderError, loading: unloadProviderLoading, reset: resetUnloadProvider },
  ] = useUnloadProviderMutation();

  useEffect(() => {
    if (getProviderError) enqueueSnackbar({ message: `Failed to get provider: ${getProviderError.message}`, variant: "error" });
    if (createProviderError) enqueueSnackbar({ message: `Failed to create provider: ${createProviderError.message}`, variant: "error" });
    if (updateProviderError) enqueueSnackbar({ message: `Failed to update provider: ${updateProviderError.message}`, variant: "error" });
    if (configureProviderError)
      enqueueSnackbar({ message: `Failed to reconfigure provider: ${configureProviderError.message}`, variant: "error" });
    if (loadProviderError) enqueueSnackbar({ message: `Failed to load provider: ${loadProviderError.message}`, variant: "error" });
    if (unloadProviderError) enqueueSnackbar({ message: `Failed to unload provider: ${unloadProviderError.message}`, variant: "error" });
  }, [getProviderError, createProviderError, updateProviderError, configureProviderError, loadProviderError, unloadProviderError]);

  // If we're editing the provider, fetch based on ID
  useEffect(() => {
    if (action === "edit" && id)
      getProvider({
        variables: {
          id,
        },
        pollInterval: 1000,
      }).catch(console.error);
  }, [id, action, getProvider]);

  // Once we get edit data, update local form state
  useEffect(() => {
    if (getProviderData?.provider)
      setProvider({
        displayName: getProviderData.provider.displayName,
        configBytes: getProviderData.provider.configBytes,
        providerGitUrl: getProviderData.provider.providerGitUrl,
        providerVersion: getProviderData.provider.providerVersion,
      } as ProviderInput);
  }, [getProviderData]);

  // Get repo tags
  useEffect(() => {
    const timeoutId = setTimeout(() => {
      if (provider.providerGitUrl) {
        const repoPath = provider.providerGitUrl.replace(/https?:\/\/github.com\//, "").replace(/\.git$/, "");
        // Check path is 2 elements
        if (repoPath.split("/").length !== 2) return;
        fetch(`https://api.github.com/repos/${repoPath}/git/refs/tags`)
          .then((res) => {
            if (res.status === 404) throw Error("repo not found");
            else return res.json();
          })
          .then((tags) => {
            if (!tags["message"]) setGitTags(tags as GitTag[]);
          })
          .catch(() => setGitTags([])); // Ignore all errors here and blank autocomplete
      }
    }, 1000);
    return () => clearTimeout(timeoutId);
  }, [provider.providerGitUrl]);

  // Move to edit page after creation
  useEffect(() => {
    if (createProviderData?.createProvider.id) {
      enqueueSnackbar({ message: "Created provider!", variant: "success" });
      navigate(`/providers/edit/${createProviderData.createProvider.id}`);
    }
  }, [createProviderData, navigate]);

  // Show pop up when updating provider
  useEffect(() => {
    if (updateProviderData?.updateProvider.id) {
      enqueueSnackbar({ message: "Updated provider!", variant: "success" });
      resetUpdateProvider();
    }
  }, [updateProviderData, enqueueSnackbar, resetUpdateProvider]);

  // Show pop up when loading provider
  useEffect(() => {
    if (loadProviderData?.loadProvider.id) {
      enqueueSnackbar({ message: "Loaded provider!", variant: "success" });
      resetLoadProvider();
    }
  }, [loadProviderData, enqueueSnackbar, resetLoadProvider]);

  // Show pop up when unloading provider
  useEffect(() => {
    if (unloadProviderData?.unloadProvider.id) {
      enqueueSnackbar({ message: "Unloaded provider!", variant: "success" });
      resetUnloadProvider();
    }
  }, [unloadProviderData, enqueueSnackbar, resetUnloadProvider]);

  // Show pop up when reconfiguring provider
  useEffect(() => {
    if (configureProviderData?.configureProvider.id) {
      enqueueSnackbar({ message: "Reconfigured provider!", variant: "success" });
      resetConfigureProvider();
    }
  }, [configureProviderData, enqueueSnackbar, resetConfigureProvider]);

  const handlePrettify = () => {
    if (monaco)
      monaco.editor.getEditors().forEach((editor) => {
        editor.getAction("editor.action.formatDocument")?.run().catch(console.error);
      });
  };

  const handleSubmitProvider = () => {
    if (action === "create")
      createProvider({
        variables: {
          input: provider,
        },
      }).catch(console.error);
    else if (action === "edit" && id)
      updateProvider({
        variables: {
          id,
          input: provider,
        },
      }).catch(console.error);
  };

  return (
    <Container sx={{ py: 3 }}>
      <Box sx={{ display: "flex", alignItems: "center", justifyContent: "space-between" }}>
        <Box sx={{ display: "flex", alignItems: "center" }}>
          <TextField
            variant="standard"
            placeholder="Untitled Provider"
            value={provider.displayName}
            onChange={(e) => setProvider((prev) => ({ ...prev, displayName: e.target.value }))}
          ></TextField>
          {getProviderData && (
            <>
              <Circle sx={{ ml: 1, height: "1rem" }} color={getProviderData.provider.isLoaded ? "success" : "error"} />
              <Typography variant="subtitle1" color={getProviderData.provider.isLoaded ? "success.main" : "error.main"}>
                {getProviderData.provider.isLoaded ? "Loaded" : "Not Loaded"}
              </Typography>
            </>
          )}
        </Box>
        <Box sx={{ display: "flex", alignItems: "center" }}>
          {getProviderData && (
            <ButtonGroup sx={{ mr: 1 }}>
              <LoadingButton
                variant="outlined"
                color="info"
                startIcon={<Cached />}
                disabled={!getProviderData.provider.isLoaded || updateProviderLoading}
                loading={configureProviderLoading}
                onClick={() => {
                  if (getProviderData.provider.isLoaded)
                    configureProvider({ variables: { id: getProviderData.provider.id } }).catch(console.error);
                }}
              >
                Reconfigure
              </LoadingButton>
              <LoadingButton
                variant="outlined"
                color={getProviderData.provider.isLoaded ? "error" : "success"}
                startIcon={getProviderData.provider.isLoaded ? <PowerOff /> : <Power />}
                loading={loadProviderLoading || unloadProviderLoading}
                disabled={updateProviderLoading}
                onClick={() => {
                  if (getProviderData.provider.isLoaded)
                    unloadProvider({ variables: { id: getProviderData.provider.id } }).catch(console.error);
                  else loadProvider({ variables: { id: getProviderData.provider.id } }).catch(console.error);
                }}
              >
                {getProviderData.provider.isLoaded ? "Unload" : "Load"}
              </LoadingButton>
            </ButtonGroup>
          )}
          <LoadingButton
            variant="contained"
            startIcon={action === "create" ? <Add /> : <Save />}
            loading={createProviderLoading || updateProviderLoading}
            disabled={
              provider.configBytes == "" ||
              provider.displayName == "" ||
              provider.providerGitUrl == "" ||
              provider.providerVersion == "" ||
              createProviderData != null
            }
            onClick={handleSubmitProvider}
          >
            {action === "create" ? "Create" : "Save"}
          </LoadingButton>
        </Box>
      </Box>
      <Divider sx={{ my: 2 }} />
      {getProviderLoading && <LinearProgress />}
      <Grid container spacing={2}>
        <Grid item md={6}>
          <TextField
            sx={{ width: "100%" }}
            label="Git URL"
            placeholder="https://github.com/cble-platform/provider-xxx"
            value={provider.providerGitUrl}
            onChange={(e) => setProvider((prev) => ({ ...prev, providerGitUrl: e.target.value }))}
          />
        </Grid>
        <Grid item md={6}>
          <Autocomplete
            freeSolo
            sx={{ width: "100%" }}
            placeholder="v1.0.0"
            options={gitTags.map((option) => option.ref.replace("refs/tags/", ""))}
            renderInput={(params) => <TextField {...params} label="Version" />}
            value={provider.providerVersion}
            onChange={(_, value) => setProvider((prev) => ({ ...prev, providerVersion: value || "" }))}
            onInputChange={(_, value) => setProvider((prev) => ({ ...prev, providerVersion: value || "" }))}
          />
        </Grid>
      </Grid>
      <Typography variant="h6" sx={{ my: 2, display: "flex", alignItems: "center" }}>
        Configuration
        <IconButton onClick={handlePrettify} sx={{ ml: 1 }} title="Prettify">
          <AutoFixHigh />
        </IconButton>
      </Typography>
      <Box sx={{ height: "25dvh", borderWidth: 1, borderStyle: "solid", borderColor: "primary" }}>
        <MonacoEditor
          theme={themePreference === "auto" ? (prefersDarkMode ? "vs-dark" : "light") : themePreference === "dark" ? "vs-dark" : "light"}
          language="yaml"
          options={{ minimap: { enabled: false }, scrollBeyondLastLine: false, wordWrap: "on" }}
          value={provider.configBytes}
          onChange={(value, _) => setProvider((prev) => ({ ...prev, configBytes: value ?? "" }))}
        ></MonacoEditor>
      </Box>
    </Container>
  );
}
