import { useSnackbar } from 'notistack'
import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import {
  DeployResourceType,
  GetResourcesQuery,
  useGetConsoleMutation,
  useGetDeploymentQuery,
  useGetResourcesQuery,
} from '@/api/generated'
import {
  Container,
  Typography,
  Divider,
  LinearProgress,
  Box,
  Menu,
  MenuItem,
  Button,
  List,
  ListItem,
  ListItemText,
  IconButton,
  CircularProgress,
  ListItemButton,
  ListItemIcon,
} from '@mui/material'
import MuiMarkdown from 'mui-markdown'
import {
  ChevronLeft,
  Computer,
  ExpandMore,
  Help,
  Lan,
  Monitor,
  MoreHoriz,
  Router,
} from '@mui/icons-material'

function getResourceTypeIcon(
  resource: GetResourcesQuery['deploymentResources'][number]
): React.ReactElement {
  switch (resource.type) {
    case DeployResourceType.Host:
      return <Computer />
    case DeployResourceType.Network:
      return <Lan />
    case DeployResourceType.Router:
      return <Router />
    case DeployResourceType.Unknown:
    default:
      return <Help />
  }
}

export default function DeploymentDetails() {
  const { id } = useParams()
  const { enqueueSnackbar } = useSnackbar()
  const navigate = useNavigate()
  const {
    data: getDeploymentData,
    error: getDeploymentError,
    loading: getDeploymentLoading,
  } = useGetDeploymentQuery({ variables: { id: id || '' } })
  const {
    data: getResourcesData,
    error: getResourcesError,
    loading: getResourcesLoading,
  } = useGetResourcesQuery({ variables: { id: id || '' } })
  const [
    getConsole,
    {
      data: getConsoleData,
      error: getConsoleError,
      loading: getConsoleLoading,
      reset: resetGetConsole,
    },
  ] = useGetConsoleMutation()
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null)
  const [resourceMoreMenuEl, setResourceMoreMenuEl] = useState<null | {
    el: HTMLElement
    resource: GetResourcesQuery['deploymentResources'][number]
  }>(null)
  const [selectedResource, setSelectedResource] = useState<
    GetResourcesQuery['deploymentResources'][number] | null
  >()

  useEffect(() => {
    if (getDeploymentError)
      enqueueSnackbar({
        message: `Failed to get deployment: ${getDeploymentError.message}`,
        variant: 'error',
      })
    if (getResourcesError)
      enqueueSnackbar({
        message: `Failed to get resources: ${getResourcesError.message}`,
        variant: 'error',
      })
  }, [getDeploymentError, getResourcesError])

  // Auto open new tab to console
  useEffect(() => {
    if (getConsoleData) {
      console.log(getConsoleData.getConsole)
      window.open(getConsoleData.getConsole, '_blank', 'noreferrer')
      resetGetConsole()
      setResourceMoreMenuEl(null)
    }
  }, [getConsoleData])

  return (
    <Container sx={{ py: 3 }}>
      <Button href="/deployments" startIcon={<ChevronLeft />} sx={{ mb: 2 }}>
        Back
      </Button>
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
        }}
      >
        <Typography variant="h4">
          Details - {getDeploymentData?.deployment.name}
        </Typography>
        <Button
          id="more-button"
          aria-controls={moreMenuEl ? 'more-menu' : undefined}
          aria-haspopup="true"
          aria-expanded={moreMenuEl ? 'true' : undefined}
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
            'aria-labelledby': 'more-button',
          }}
        >
          <MenuItem
            onClick={() =>
              navigate(
                `/deployments/destroy/${getDeploymentData?.deployment.id}`
              )
            }
          >
            Destroy
          </MenuItem>
        </Menu>
      </Box>
      <Divider sx={{ my: 2 }} />
      {getDeploymentLoading && <LinearProgress sx={{ my: 2 }} />}
      <Box
        sx={{
          display: 'grid',
          gridTemplateColumns: '1fr 3fr',
          gridTemplateRows: '1fr',
          gap: 2,
        }}
      >
        <List dense sx={{ borderRight: '1px solid', pr: 1 }}>
          {getResourcesLoading && (
            <ListItem>
              <LinearProgress sx={{ width: '100%', mx: 1 }} />
            </ListItem>
          )}
          {getResourcesData?.deploymentResources.map((resource) => (
            <ListItem
              key={resource.key}
              secondaryAction={
                <IconButton
                  size="small"
                  edge="end"
                  aria-label="more"
                  id="resource-more-button"
                  aria-controls={resourceMoreMenuEl ? 'more-menu' : undefined}
                  aria-haspopup="true"
                  aria-expanded={resourceMoreMenuEl ? 'true' : undefined}
                  onClick={(e) =>
                    setResourceMoreMenuEl({
                      resource,
                      el: e.currentTarget,
                    })
                  }
                  sx={{ m: 0 }}
                >
                  <MoreHoriz />
                </IconButton>
              }
              sx={{ px: 1, py: 0 }}
            >
              <ListItemButton
                selected={selectedResource?.key === resource.key}
                onClick={() =>
                  setSelectedResource(
                    selectedResource?.key === resource.key ? null : resource
                  )
                }
              >
                <ListItemIcon sx={{ minWidth: '2.5rem' }}>
                  {getResourceTypeIcon(resource)}
                </ListItemIcon>
                <ListItemText primary={resource.name} />
              </ListItemButton>
            </ListItem>
          ))}
          <Menu
            id="resource-more-menu"
            anchorEl={resourceMoreMenuEl?.el}
            open={Boolean(resourceMoreMenuEl?.el)}
            onClose={() => setResourceMoreMenuEl(null)}
            MenuListProps={{
              'aria-labelledby': 'resource-more-button',
            }}
          >
            {resourceMoreMenuEl?.resource.type === DeployResourceType.Host && (
              <MenuItem
                onClick={() =>
                  getConsole({
                    variables: {
                      id: id ?? '',
                      hostKey: resourceMoreMenuEl?.resource.key ?? '',
                    },
                  })
                }
                disabled={getConsoleLoading}
              >
                {getConsoleLoading ? (
                  <CircularProgress sx={{ mr: 1 }} size="1rem" />
                ) : (
                  <Monitor sx={{ mr: 1 }} />
                )}
                Console
              </MenuItem>
            )}
          </Menu>
        </List>
        <Box sx={{ px: 3, py: 1 }}>
          {selectedResource ? (
            <Box>
              <Typography variant="h4">{selectedResource.name}</Typography>
            </Box>
          ) : (
            <MuiMarkdown>
              {getDeploymentData?.deployment.blueprint.description}
            </MuiMarkdown>
          )}
        </Box>
      </Box>
    </Container>
  )
}
