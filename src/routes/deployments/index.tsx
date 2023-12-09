import { Container, Box, Typography, Divider, Grid, CardContent, CardHeader, Card, IconButton, Menu, MenuItem } from "@mui/material";
import { useDestroyDeploymentMutation, useListDeploymentsQuery, useUpdateDeploymentMutation } from "../../api/graphql/generated";
import { useSnackbar } from "notistack";
import { useEffect, useState } from "react";
import { MoreVert } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";

export default function Deployments() {
  const navigate = useNavigate();
  const { enqueueSnackbar } = useSnackbar();
  const { data: listDeploymentsData, error: listDeploymentsError, loading: listDeploymentsLoading } = useListDeploymentsQuery();
  const [
    updateDeployment,
    { data: updateDeploymentData, error: updateDeploymentError, loading: updateDeploymentLoading, reset: resetUpdateDeployment },
  ] = useUpdateDeploymentMutation();
  const [
    destroyDeployment,
    { data: destroyDeploymentData, error: destroyDeploymentError, loading: destroyDeploymentLoading, reset: resetDestroyDeployment },
  ] = useDestroyDeploymentMutation();
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null);
  const [moreMenuId, setMoreMenuId] = useState<null | string>(null);

  useEffect(() => {
    if (listDeploymentsError) enqueueSnackbar({ message: `Failed to get deployments: ${listDeploymentsError.message}`, variant: "error" });
    if (updateDeploymentError)
      enqueueSnackbar({ message: `Failed to update deployment: ${updateDeploymentError.message}`, variant: "error" });
    if (destroyDeploymentError)
      enqueueSnackbar({ message: `Failed to destroy deployment: ${destroyDeploymentError.message}`, variant: "error" });
  }, [listDeploymentsError, updateDeploymentError, destroyDeploymentError]);

  const handleMoreMenuClick = (id: string, event: React.MouseEvent<HTMLElement>) => {
    setMoreMenuEl(event.currentTarget);
    setMoreMenuId(id);
  };
  const handleMoreMenuClose = () => {
    setMoreMenuEl(null);
    setMoreMenuId(null);
  };

  return (
    <Container sx={{ py: 3 }}>
      <Box sx={{ display: "flex", alignContent: "center", justifyContent: "space-between" }}>
        <Typography variant="h4">Deployments</Typography>
      </Box>
      <Divider sx={{ my: 3 }} />
      <Grid container spacing={2}>
        {listDeploymentsData?.deployments.map((deployment) => (
          <Grid item xs={12} key={deployment.id}>
            <Card sx={{ width: "100%" }}>
              <CardHeader
                title={deployment.blueprint.name}
                action={
                  <IconButton
                    id="more-button"
                    aria-controls={moreMenuEl ? "more-menu" : undefined}
                    aria-haspopup="true"
                    aria-expanded={moreMenuEl ? "true" : undefined}
                    onClick={(event) => handleMoreMenuClick(deployment.id as string, event)}
                  >
                    <MoreVert />
                  </IconButton>
                }
              />
              <CardContent>This is content</CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
      <Menu
        id="more-menu"
        anchorEl={moreMenuEl}
        open={Boolean(moreMenuEl)}
        onClose={handleMoreMenuClose}
        MenuListProps={{
          "aria-labelledby": "more-button",
        }}
      >
        <MenuItem onClick={() => navigate(`/deployments/destroy/${moreMenuId}`)}>Destroy</MenuItem>
        <MenuItem onClick={() => navigate(`/deployments/${moreMenuId}`)}>Details</MenuItem>
      </Menu>
    </Container>
  );
}
