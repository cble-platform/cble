import { useSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useGetDeploymentLazyQuery } from "../../api/graphql/generated";
import { Container, Typography, Divider, LinearProgress, Box, Menu, MenuItem, Button } from "@mui/material";
import MuiMarkdown from "mui-markdown";
import { ChevronLeft, ExpandMore } from "@mui/icons-material";

export default function DeploymentDetails() {
  const { id } = useParams();
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const [getDeployment, { data: getDeploymentData, error: getDeploymentError, loading: getDeploymentLoading }] =
    useGetDeploymentLazyQuery();
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null);

  useEffect(() => {
    if (id) getDeployment({ variables: { id } });
  }, [id]);

  useEffect(() => {
    if (getDeploymentError) enqueueSnackbar({ message: `Failed to get deployment: ${getDeploymentError.message}`, variant: "error" });
  }, [getDeploymentError]);

  return (
    <Container sx={{ py: 3 }}>
      <Button href="/deployments" startIcon={<ChevronLeft />} sx={{ mb: 2 }}>
        Back
      </Button>
      <Box sx={{ display: "flex", alignItems: "center", justifyContent: "space-between" }}>
        <Typography variant="h4">Details - {getDeploymentData?.deployment.name}</Typography>
        <Button
          id="more-button"
          aria-controls={moreMenuEl ? "more-menu" : undefined}
          aria-haspopup="true"
          aria-expanded={moreMenuEl ? "true" : undefined}
          onClick={(e) => setMoreMenuEl(e.currentTarget)}
          startIcon={<ExpandMore />}
        >
          Actions
        </Button>
        <Menu
          id="more-menu"
          anchorEl={moreMenuEl}
          open={Boolean(moreMenuEl)}
          onClose={() => setMoreMenuEl(null)}
          MenuListProps={{
            "aria-labelledby": "more-button",
          }}
        >
          <MenuItem onClick={() => navigate(`/deployments/destroy/${getDeploymentData?.deployment.id}`)}>Destroy</MenuItem>
        </Menu>
      </Box>
      <Divider sx={{ my: 2 }} />
      {getDeploymentLoading && <LinearProgress sx={{ my: 2 }} />}
      <MuiMarkdown>{getDeploymentData?.deployment.blueprint.description}</MuiMarkdown>
    </Container>
  );
}
