import {
  Container,
  Box,
  Card,
  CardContent,
  Typography,
  Divider,
  Grid,
  CardActions,
  Button,
  TypographyProps,
  IconButton,
  LinearProgress,
} from "@mui/material";
import { useBlueprintsQuery, useMeHasPermissionQuery } from "../../api/graphql/generated";
import { useEffect } from "react";
import { MuiMarkdown, getOverrides } from "mui-markdown";
import { useSnackbar } from "notistack";
import { Edit } from "@mui/icons-material";

const MarkdownOverrides = {
  ...getOverrides(),
  h1: {
    component: Typography,
    props: {
      variant: "h6",
    } as TypographyProps,
  },
  h2: {
    component: Typography,
    props: {
      variant: "h6",
    } as TypographyProps,
  },
  h3: {
    component: Typography,
    props: {
      variant: "h6",
    } as TypographyProps,
  },
  h4: {
    component: Typography,
    props: {
      variant: "subtitle1",
    } as TypographyProps,
  },
  h5: {
    component: Typography,
    props: {
      variant: "subtitle2",
    } as TypographyProps,
  },
  h6: {
    component: Typography,
    props: {
      variant: "subtitle2",
    } as TypographyProps,
  },
};

export default function Blueprints() {
  const { data: blueprintsData, error: blueprintsError, loading: blueprintsLoading } = useBlueprintsQuery();
  const { data: createPermData } = useMeHasPermissionQuery({ variables: { key: "com.cble.blueprints.create" } });
  const { enqueueSnackbar } = useSnackbar();

  useEffect(() => {
    if (blueprintsError)
      enqueueSnackbar({
        message: `Failed to get blueprints: ${blueprintsError.message}`,
        variant: "error",
      });
  }, [blueprintsError]);

  return (
    <Container sx={{ py: 3 }}>
      <Box sx={{ display: "flex", alignContent: "center", justifyContent: "space-between" }}>
        <Typography variant="h4">Blueprints</Typography>
        {createPermData?.meHasPermission && (
          <Button href="/blueprints/create" variant="contained" color="primary">
            Create
          </Button>
        )}
      </Box>
      <Divider sx={{ my: 3 }} />
      <Grid container spacing={2}>
        {blueprintsLoading && (
          <Grid item xs={12}>
            <LinearProgress />
          </Grid>
        )}
        {blueprintsData?.blueprints.map((blueprint) => (
          <Grid item xs={3} key={blueprint.id}>
            <Card sx={{ height: "100%" }}>
              <CardContent>
                <Typography variant="h6">{blueprint.name}</Typography>
                <Typography variant="subtitle1">Group: {blueprint.parentGroup.name}</Typography>
                <Divider sx={{ my: 1 }} />
                <MuiMarkdown overrides={MarkdownOverrides}>{blueprint.description}</MuiMarkdown>
              </CardContent>
              <CardActions sx={{ justifyContent: "space-between" }}>
                <Button color="primary" href={`/blueprints/request/${blueprint.id}`}>
                  Request
                </Button>
                <IconButton href={`/blueprints/edit/${blueprint.id}`}>
                  <Edit />
                </IconButton>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
}
