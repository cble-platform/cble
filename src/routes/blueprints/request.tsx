import { Box, Container, Divider, Fab, Typography } from "@mui/material";
import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { useDeployBlueprintMutation, useGetBlueprintLazyQuery } from "../../api/graphql/generated";
import { MuiMarkdown } from "mui-markdown";
import { Send } from "@mui/icons-material";
import { useSnackbar } from "notistack";

export default function RequestBlueprint() {
  const { id } = useParams();
  const [getBlueprint, { data: blueprintData, error: blueprintError, loading: blueprintLoading }] = useGetBlueprintLazyQuery();
  const [deployBlueprint, { data: deployBlueprintData, error: deployBlueprintError, loading: deployBlueprintLoading }] =
    useDeployBlueprintMutation();
  const { enqueueSnackbar } = useSnackbar();

  useEffect(() => {
    // Isn't loading, isn't already loaded, isn't errored, and has id
    if (!blueprintLoading && !blueprintData && !blueprintError && id)
      getBlueprint({
        variables: {
          id,
        },
      });
  }, [id]);

  useEffect(() => {
    if (deployBlueprintLoading) enqueueSnackbar({ message: "Deploying blueprint...", variant: "info" });
    if (deployBlueprintError) enqueueSnackbar({ message: `Error deploying blueprint: ${deployBlueprintError.message}`, variant: "error" });
    if (deployBlueprintData) enqueueSnackbar({ message: "Deployed blueprint!", variant: "success" });
  }, [deployBlueprintData, deployBlueprintError, deployBlueprintLoading]);

  return (
    <Container sx={{ display: "flex", flexDirection: "column", height: "100%", py: 2 }}>
      <Box sx={{ width: "100%", display: "flex", alignItems: "center" }}>
        <Typography variant="h4">Request {blueprintData?.blueprint.name}</Typography>
      </Box>
      <Divider sx={{ width: "100%", my: 2 }} />
      <MuiMarkdown>{blueprintData?.blueprint.description}</MuiMarkdown>
      <Fab
        color="primary"
        variant="extended"
        sx={{ position: "fixed", bottom: "2rem", right: "2rem" }}
        onClick={() => {
          if (id)
            deployBlueprint({
              variables: {
                id,
              },
            });
        }}
        disabled={!id || deployBlueprintLoading || deployBlueprintData != null || deployBlueprintError != null}
      >
        <Send sx={{ mr: 1 }} /> Request
      </Fab>
    </Container>
  );
}
