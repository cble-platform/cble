import {
  Container,
  Box,
  Typography,
  Divider,
  Grid,
  CardContent,
  Card,
  IconButton,
  Menu,
  MenuItem,
  TextField,
  CircularProgress,
  LinearProgress,
  Button,
  Link,
} from '@mui/material'
import {
  useListDeploymentsQuery,
  useUpdateDeploymentMutation,
} from '../../api/generated'
import { useSnackbar } from 'notistack'
import React, { useEffect, useState } from 'react'
import { Cancel, Edit, ExpandMore, Save } from '@mui/icons-material'
import { useNavigate } from 'react-router-dom'

export default function Deployments() {
  const navigate = useNavigate()
  const { enqueueSnackbar } = useSnackbar()
  const {
    data: listDeploymentsData,
    error: listDeploymentsError,
    loading: listDeploymentsLoading,
    refetch: refetchListDeployments,
  } = useListDeploymentsQuery()
  const [
    updateDeployment,
    {
      data: updateDeploymentData,
      error: updateDeploymentError,
      loading: updateDeploymentLoading,
      reset: resetUpdateDeployment,
    },
  ] = useUpdateDeploymentMutation()
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null)
  const [moreMenuId, setMoreMenuId] = useState<null | string>(null)
  const [editDeploymentNameData, setEditDeploymentNameData] = useState<null | {
    id: string
    name: string
  }>(null)

  useEffect(() => {
    if (listDeploymentsError)
      enqueueSnackbar({
        message: `Failed to get deployments: ${listDeploymentsError.message}`,
        variant: 'error',
      })
    if (updateDeploymentError)
      enqueueSnackbar({
        message: `Failed to update deployment: ${updateDeploymentError.message}`,
        variant: 'error',
      })
  }, [listDeploymentsError, updateDeploymentError])

  useEffect(() => {
    if (updateDeploymentData) {
      enqueueSnackbar({ message: 'Updated deployment!', variant: 'success' })
      resetUpdateDeployment()
      setEditDeploymentNameData(null)
      refetchListDeployments().catch(console.error)
    }
  }, [updateDeploymentData, enqueueSnackbar, resetUpdateDeployment])

  const handleMoreMenuClick = (
    id: string,
    event: React.MouseEvent<HTMLElement>
  ) => {
    setMoreMenuEl(event.currentTarget)
    setMoreMenuId(id)
  }
  const handleMoreMenuClose = () => {
    setMoreMenuEl(null)
    setMoreMenuId(null)
  }
  const handleSaveDeployment = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (editDeploymentNameData)
      updateDeployment({
        variables: {
          id: editDeploymentNameData.id,
          input: {
            name: editDeploymentNameData.name,
          },
        },
      }).catch(console.error)
  }

  return (
    <Container sx={{ py: 3 }}>
      <Box
        sx={{
          display: 'flex',
          alignContent: 'center',
          justifyContent: 'space-between',
        }}
      >
        <Typography variant="h4">Deployments</Typography>
      </Box>
      <Divider sx={{ my: 3 }} />
      <Grid container spacing={2}>
        {listDeploymentsLoading && (
          <Grid item xs={12} sx={{ my: 2 }}>
            <LinearProgress />
          </Grid>
        )}
        {listDeploymentsData?.deployments.map((deployment) => {
          const createdDaysDiff = Math.ceil(
            (Date.now() - new Date(deployment.createdAt as string).getTime()) /
              1000 /
              60 /
              60 /
              24
          )
          return (
            <Grid item xs={12} key={deployment.id}>
              <Card sx={{ width: '100%' }}>
                <CardContent>
                  <Grid
                    container
                    spacing={1}
                    sx={{
                      '& .MuiGrid-item': {
                        display: 'flex',
                        alignItems: 'center',
                      },
                    }}
                  >
                    <Grid item xs={4}>
                      <Box
                        sx={{
                          display: 'flex',
                          alignItems: 'center',
                          '&:hover .MuiIconButton-root': {
                            visibility: 'visible',
                          },
                        }}
                      >
                        {editDeploymentNameData?.id === deployment.id ? (
                          <form onSubmit={handleSaveDeployment}>
                            <TextField
                              variant="standard"
                              value={editDeploymentNameData.name}
                              onChange={(e) =>
                                setEditDeploymentNameData({
                                  id: deployment.id,
                                  name: e.target.value ?? '',
                                })
                              }
                              disabled={updateDeploymentLoading}
                            />
                            <IconButton
                              sx={{ ml: 1 }}
                              size="small"
                              type="submit"
                              disabled={updateDeploymentLoading}
                            >
                              {updateDeploymentLoading ? (
                                <CircularProgress size="1rem" />
                              ) : (
                                <Save />
                              )}
                            </IconButton>
                            <IconButton
                              sx={{ ml: 1 }}
                              size="small"
                              onClick={() => setEditDeploymentNameData(null)}
                              disabled={updateDeploymentLoading}
                            >
                              <Cancel />
                            </IconButton>
                          </form>
                        ) : (
                          <>
                            <Link
                              href={`/deployments/${deployment.id}`}
                              variant="h5"
                              color="primary"
                            >
                              {deployment.name}
                            </Link>
                            <IconButton
                              sx={{ ml: 1, visibility: 'hidden' }}
                              size="small"
                              onClick={() =>
                                setEditDeploymentNameData({
                                  id: deployment.id,
                                  name: deployment.name,
                                })
                              }
                            >
                              <Edit />
                            </IconButton>
                          </>
                        )}
                      </Box>
                    </Grid>
                    <Grid item xs={4}>
                      <Typography variant="subtitle1">
                        Created {createdDaysDiff} day
                        {createdDaysDiff > 1 ? 's' : ''} ago
                      </Typography>
                    </Grid>
                    <Grid
                      item
                      xs={4}
                      sx={{ display: 'flex', justifyContent: 'flex-end' }}
                    >
                      <Button
                        id="more-button"
                        aria-controls={moreMenuEl ? 'more-menu' : undefined}
                        aria-haspopup="true"
                        aria-expanded={moreMenuEl ? 'true' : undefined}
                        onClick={(event) =>
                          handleMoreMenuClick(deployment.id, event)
                        }
                        startIcon={<ExpandMore />}
                      >
                        Actions
                      </Button>
                    </Grid>
                    <Grid item xs={4}>
                      <Typography variant="body1">
                        Owner: {deployment.requester.firstName}{' '}
                        {deployment.requester.lastName}
                        {/* <br />
                        Group: {deployment.blueprint.parentGroup.name} */}
                      </Typography>
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          )
        })}
      </Grid>
      <Menu
        id="more-menu"
        anchorEl={moreMenuEl}
        open={Boolean(moreMenuEl)}
        onClose={handleMoreMenuClose}
        MenuListProps={{
          'aria-labelledby': 'more-button',
        }}
      >
        <MenuItem
          onClick={() => navigate(`/deployments/destroy/${moreMenuId}`)}
        >
          Destroy
        </MenuItem>
        <MenuItem onClick={() => navigate(`/deployments/${moreMenuId}`)}>
          Details
        </MenuItem>
      </Menu>
    </Container>
  )
}
