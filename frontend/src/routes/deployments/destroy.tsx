import {
  Button,
  Container,
  Divider,
  LinearProgress,
  Typography,
  styled,
} from '@mui/material'
import { useNavigate, useParams } from 'react-router-dom'
import {
  useDestroyDeploymentMutation,
  useGetDeploymentLazyQuery,
} from '../../lib/api/generated'
import { useEffect } from 'react'
import { useSnackbar } from 'notistack'
import { ChevronLeft, Delete } from '@mui/icons-material'
import ContainerFab from '../../components/container-fab'

const ErrorSpan = styled('span')(({ theme }) => ({
  color: theme.palette.error.main,
  fontWeight: 'bold',
}))

export default function DestroyDeployment() {
  const { id } = useParams()
  const { enqueueSnackbar } = useSnackbar()
  const navigate = useNavigate()
  const [
    getDeployment,
    {
      data: getDeploymentData,
      error: getDeploymentError,
      loading: getDeploymentLoading,
    },
  ] = useGetDeploymentLazyQuery()
  const [
    destroyDeployment,
    {
      data: destroyDeploymentData,
      error: destroyDeploymentError,
      loading: destroyDeploymentLoading,
    },
  ] = useDestroyDeploymentMutation()

  useEffect(() => {
    if (id) getDeployment({ variables: { id } }).catch(console.error)
  }, [id])

  useEffect(() => {
    if (getDeploymentError)
      enqueueSnackbar({
        message: `Failed to get deployment: ${getDeploymentError.message}`,
        variant: 'error',
      })
    if (destroyDeploymentError)
      enqueueSnackbar({
        message: `Failed to destroy deployment: ${destroyDeploymentError.message}`,
        variant: 'error',
      })
  }, [getDeploymentError, destroyDeploymentError])

  useEffect(() => {
    if (destroyDeploymentData?.destroyDeployment) {
      enqueueSnackbar({
        message: `Destroyed deployment "${getDeploymentData?.deployment.name}"!`,
        variant: 'success',
      })
      navigate('/deployments')
    }
  })

  return (
    <Container sx={{ py: 3 }}>
      <Button href="/deployments" startIcon={<ChevronLeft />} sx={{ mb: 2 }}>
        Back
      </Button>
      <Typography variant="h4">
        Destroy - {getDeploymentData?.deployment.name}
      </Typography>
      <Divider sx={{ my: 2 }} />
      <Typography variant="body1">
        This will destroy this deployment and all resources created as a part of
        the deployment. <ErrorSpan>All data will be lost!</ErrorSpan>
      </Typography>
      {getDeploymentLoading && <LinearProgress sx={{ my: 2 }} />}
      <ContainerFab
        variant="extended"
        color="primary"
        onClick={() => {
          if (id) destroyDeployment({ variables: { id } }).catch(console.error)
        }}
        disabled={
          getDeploymentData === undefined ||
          destroyDeploymentLoading ||
          destroyDeploymentData != null
        }
      >
        <Delete sx={{ mr: 1 }} /> Destroy
      </ContainerFab>
    </Container>
  )
}
