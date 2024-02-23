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
  Stack,
  Checkbox,
  FormControlLabel,
} from '@mui/material'
import {
  DeploymentState,
  ListMyDeploymentsQuery,
  useListMyDeploymentsQuery,
  useUpdateDeploymentMutation,
} from '../../lib/api/generated'
import { useSnackbar } from 'notistack'
import React, { useEffect, useState } from 'react'
import { Cancel, Edit, ExpandMore, Save } from '@mui/icons-material'
import { useNavigate } from 'react-router-dom'

function daysBetween(
  date1: string | number | Date,
  date2: string | number | Date
) {
  return (
    (new Date(date1).getTime() - new Date(date2).getTime()) /
    1000 /
    60 /
    60 /
    24
  )
}

function generateCreatedMessage(
  deployment: ListMyDeploymentsQuery['myDeployments']['deployments'][number]
): string {
  const createdDaysDiff = daysBetween(
    new Date(deployment.createdAt as string),
    Date.now()
  )
  if (createdDaysDiff < 1) return 'Created Today'
  if (createdDaysDiff < 2) return 'Created Yesterday'
  if (createdDaysDiff > 30)
    return `Created on ${new Date(deployment.createdAt as string).toDateString()}`
  else return `${Math.floor(createdDaysDiff)} days ago`
}

function generateExpiryMessage(
  deployment: ListMyDeploymentsQuery['myDeployments']['deployments'][number]
): string {
  const expiresDaysDiff = daysBetween(
    new Date(deployment.expiresAt as string),
    Date.now()
  )
  let expireTense = ''
  if (expiresDaysDiff < 0) expireTense = 'Expired'
  else expireTense = 'Expires'
  let dayDisplay
  if (Math.abs(expiresDaysDiff) < 1) dayDisplay = 'Today'
  else if (expiresDaysDiff < -2)
    dayDisplay = `days ${Math.ceil(Math.abs(expiresDaysDiff))} ago`
  else if (expiresDaysDiff < -1) dayDisplay = 'Yesterday'
  else if (expiresDaysDiff < 2) dayDisplay = 'Tomorrow'
  else dayDisplay = `in ${Math.floor(Math.abs(expiresDaysDiff))} days`

  return `${expireTense} ${dayDisplay}`
}

export default function Deployments() {
  const navigate = useNavigate()
  const { enqueueSnackbar } = useSnackbar()
  const [includeExpiredAndDestroyed, setIncludeExpired] =
    useState<boolean>(false)
  const {
    data: listMyDeploymentsData,
    error: listMyDeploymentsError,
    loading: listMyDeploymentsLoading,
    refetch: refetchListMyDeployments,
  } = useListMyDeploymentsQuery({
    variables: {
      includeExpiredAndDestroyed,
    },
  })
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
    if (listMyDeploymentsError)
      enqueueSnackbar({
        message: `Failed to get deployments: ${listMyDeploymentsError.message}`,
        variant: 'error',
      })
    if (updateDeploymentError)
      enqueueSnackbar({
        message: `Failed to update deployment: ${updateDeploymentError.message}`,
        variant: 'error',
      })
  }, [listMyDeploymentsError, updateDeploymentError])

  useEffect(() => {
    if (updateDeploymentData) {
      enqueueSnackbar({ message: 'Updated deployment!', variant: 'success' })
      resetUpdateDeployment()
      setEditDeploymentNameData(null)
      refetchListMyDeployments({
        includeExpiredAndDestroyed,
      }).catch(console.error)
    }
  }, [updateDeploymentData, enqueueSnackbar, resetUpdateDeployment])

  useEffect(() => {
    refetchListMyDeployments({
      includeExpiredAndDestroyed,
    }).catch(console.error)
  }, [includeExpiredAndDestroyed])

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
      <Stack direction="row" alignContent="center" gap={4}>
        <Typography variant="h4">Deployments</Typography>
        <FormControlLabel
          control={
            <Checkbox
              checked={includeExpiredAndDestroyed}
              onChange={(_, checked) => setIncludeExpired(checked)}
            />
          }
          label="Show Expired/Destroyed"
        />
      </Stack>
      <Divider sx={{ my: 3 }} />
      <Grid container spacing={2}>
        {listMyDeploymentsLoading && (
          <Grid item xs={12} sx={{ my: 2 }}>
            <LinearProgress />
          </Grid>
        )}
        {listMyDeploymentsData?.myDeployments.deployments.map((deployment) => {
          const createdDaysDiff = daysBetween(
            new Date(deployment.createdAt as string),
            Date.now()
          )
          return (
            <Grid item xs={12} key={deployment.id}>
              <Card
                sx={{
                  width: '100%',
                }}
              >
                <CardContent
                  sx={{
                    opacity:
                      deployment.state === DeploymentState.Destroyed ? 0.5 : 1,
                    pointerEvents:
                      deployment.state === DeploymentState.Destroyed
                        ? 'none'
                        : 'auto',
                  }}
                >
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
                        {generateCreatedMessage(deployment)}
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
                      </Typography>
                    </Grid>
                    <Grid item xs={4}>
                      <Typography
                        variant="body2"
                        sx={{ color: 'text.secondary' }}
                      >
                        {generateExpiryMessage(deployment)}
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
