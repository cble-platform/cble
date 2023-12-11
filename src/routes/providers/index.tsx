import {
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Container,
  Divider,
  Grid,
  LinearProgress,
  Menu,
  MenuItem,
  Typography,
} from "@mui/material";
import { useSnackbar } from "notistack";
import {
  ListProvidersQuery,
  useListProvidersQuery,
  useLoadProviderMutation,
  useMeHasPermissionQuery,
  useUnloadProviderMutation,
} from "../../api/graphql/generated";
import { useEffect, useState } from "react";
import { TypographyCode } from "../../components/custom-typography";
import { Add, Circle, ExpandMore } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";

export default function Providers() {
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const { data: createPermData } = useMeHasPermissionQuery({ variables: { key: "com.cble.providers.create" } });
  const {
    data: listProvidersData,
    error: listProvidersError,
    loading: listProvidersLoading,
    refetch: refetchListProviders,
  } = useListProvidersQuery();
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null);
  const [moreMenuProvider, setMoreMenuProvider] = useState<ListProvidersQuery["providers"][number]>();
  const [loadProvider, { data: loadProviderData, error: loadProviderError, loading: loadProviderLoading, reset: resetLoadProvider }] =
    useLoadProviderMutation();
  const [
    unloadProvider,
    { data: unloadProviderData, error: unloadProviderError, loading: unloadProviderLoading, reset: resetUnloadProvider },
  ] = useUnloadProviderMutation();

  useEffect(() => {
    if (listProvidersError) enqueueSnackbar({ message: `Failed to list providers: ${listProvidersError.message}`, variant: "error" });
    if (loadProviderError) enqueueSnackbar({ message: `Failed to load provider: ${loadProviderError.message}`, variant: "error" });
    if (unloadProviderError) enqueueSnackbar({ message: `Failed to unload providers: ${unloadProviderError.message}`, variant: "error" });
  }, [listProvidersError, enqueueSnackbar]);

  useEffect(() => {
    if (loadProviderData) {
      enqueueSnackbar({ message: `Loaded provider "${loadProviderData.loadProvider.displayName}"`, variant: "success" });
      resetLoadProvider();
      refetchListProviders();
      handleMoreMenuClose();
    }
    if (unloadProviderData) {
      enqueueSnackbar({ message: `Unloaded provider "${unloadProviderData.unloadProvider.displayName}"`, variant: "success" });
      resetUnloadProvider();
      refetchListProviders();
      handleMoreMenuClose();
    }
  }, [loadProviderData, unloadProviderData, enqueueSnackbar, resetLoadProvider, resetUnloadProvider, refetchListProviders]);

  const handleMoreMenuClick = (provider: ListProvidersQuery["providers"][number], event: React.MouseEvent<HTMLElement>) => {
    setMoreMenuEl(event.currentTarget);
    setMoreMenuProvider(provider);
  };
  const handleMoreMenuClose = () => {
    setMoreMenuEl(null);
    // Don't clear moreMenuProvider otherwise the menu load/unload text will flash
  };

  return (
    <Container sx={{ py: 3 }}>
      <Box sx={{ display: "flex", alignContent: "center", justifyContent: "space-between" }}>
        <Typography variant="h4">Providers</Typography>
        {createPermData?.meHasPermission && (
          <Button href="/providers/create" variant="contained" color="primary" startIcon={<Add />}>
            Create
          </Button>
        )}
      </Box>
      <Divider sx={{ my: 3 }} />
      <Grid container spacing={2}>
        {listProvidersLoading && (
          <Grid item xs={12} sx={{ my: 2 }}>
            <LinearProgress />
          </Grid>
        )}
        {listProvidersData?.providers.map((provider) => (
          <Grid item xs={12} key={provider.id}>
            <Card sx={{ width: "100%" }}>
              <CardContent>
                <Grid container spacing={1}>
                  <Grid item xs={6} sx={{ display: "flex", alignItems: "center" }}>
                    <Typography variant="h5">{provider.displayName}</Typography>
                    <Circle color={provider.isLoaded ? "success" : "error"} sx={{ ml: 1, height: "1rem" }} />
                    <Typography variant="subtitle1" color={provider.isLoaded ? "success.main" : "error.main"}>
                      {provider.isLoaded ? "Loaded" : "Not Loaded"}
                    </Typography>
                  </Grid>
                  <Grid item xs={6} sx={{ display: "flex", alignItems: "center", justifyContent: "flex-end" }}>
                    <Button
                      id="more-button"
                      aria-controls={moreMenuEl ? "more-menu" : undefined}
                      aria-haspopup="true"
                      aria-expanded={moreMenuEl ? "true" : undefined}
                      onClick={(event) => handleMoreMenuClick(provider, event)}
                      startIcon={<ExpandMore />}
                    >
                      Actions
                    </Button>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography variant="body1">
                      Git Source: <TypographyCode>{provider.providerGitUrl}</TypographyCode>
                    </Typography>
                    <Typography variant="body1">
                      Version: <TypographyCode>{provider.providerVersion}</TypographyCode>
                    </Typography>
                  </Grid>
                </Grid>
              </CardContent>
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
        <MenuItem
          onClick={() => {
            if (!moreMenuProvider) return;
            if (moreMenuProvider.isLoaded) unloadProvider({ variables: { id: moreMenuProvider.id } }).catch(console.error);
            else loadProvider({ variables: { id: moreMenuProvider.id } }).catch(console.error);
          }}
          disabled={loadProviderLoading || unloadProviderLoading}
        >
          {loadProviderLoading || unloadProviderLoading ? <CircularProgress size="1rem" /> : moreMenuProvider?.isLoaded ? "Unload" : "Load"}
        </MenuItem>
        <MenuItem onClick={() => navigate(`/providers/edit/${moreMenuProvider?.id}`)}>Edit</MenuItem>
      </Menu>
    </Container>
  );
}
