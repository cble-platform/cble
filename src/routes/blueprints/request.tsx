import {
  Box,
  Container,
  Divider,
  Paper,
  Stack,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TextField,
  Typography,
} from '@mui/material'
import { useEffect, useMemo, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import {
  useDeployBlueprintMutation,
  useGetBlueprintLazyQuery,
} from '../../lib/api/generated'
import { Send } from '@mui/icons-material'
import { useSnackbar } from 'notistack'
import ContainerFab from '../../components/container-fab'
import ProjectAutocomplete, {
  ProjectAutocompleteOption,
} from '@/components/project-autocomplete'

export default function RequestBlueprint() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [
    getBlueprint,
    { data: blueprintData, error: blueprintError, loading: blueprintLoading },
  ] = useGetBlueprintLazyQuery()
  const [deployInput, setDeployInput] = useState<{
    blueprintId: string
    projectId: string
    templateVars: Record<string, string | number>
  }>({
    blueprintId: id ?? '',
    projectId: '',
    templateVars: {},
  })
  const [selectedProject, setSelectedProject] =
    useState<ProjectAutocompleteOption | null>(null)
  // Template Variables
  const [templateVars, setTemplateVars] = useState<
    readonly { name: string; value: string | number }[]
  >([])
  const [
    deployBlueprint,
    {
      data: deployBlueprintData,
      error: deployBlueprintError,
      loading: deployBlueprintLoading,
    },
  ] = useDeployBlueprintMutation()
  const { enqueueSnackbar } = useSnackbar()

  // If exceeds any quota
  const exceedsAnyQuota = useMemo<boolean>(() => {
    if (!selectedProject) return false
    if (
      selectedProject.usageCpu +
        (blueprintData?.blueprint.resources.reduce(
          (a, b) => a + b.quotaRequirements.cpu,
          0
        ) ?? 0) >
      selectedProject.quotaCpu
    )
      return true
    if (
      selectedProject.usageRam +
        (blueprintData?.blueprint.resources.reduce(
          (a, b) => a + b.quotaRequirements.ram,
          0
        ) ?? 0) >
      selectedProject.quotaRam
    )
      return true
    if (
      selectedProject.usageDisk +
        (blueprintData?.blueprint.resources.reduce(
          (a, b) => a + b.quotaRequirements.disk,
          0
        ) ?? 0) >
      selectedProject.quotaDisk
    )
      return true
    if (
      selectedProject.usageNetwork +
        (blueprintData?.blueprint.resources.reduce(
          (a, b) => a + b.quotaRequirements.network,
          0
        ) ?? 0) >
      selectedProject.quotaNetwork
    )
      return true
    if (
      selectedProject.usageRouter +
        (blueprintData?.blueprint.resources.reduce(
          (a, b) => a + b.quotaRequirements.router,
          0
        ) ?? 0) >
      selectedProject.quotaRouter
    )
      return true
    return false
  }, [blueprintData, selectedProject])

  // Pre-load values into the variables
  useEffect(() => {
    if (blueprintData?.blueprint.variableTypes)
      setTemplateVars(() =>
        Object.keys(
          blueprintData.blueprint.variableTypes as Record<
            string,
            string | number
          >
        ).map((name) => ({
          name,
          value: blueprintData.blueprint.variableTypes[name] === 'INT' ? 0 : '',
        }))
      )
  }, [blueprintData])

  // Update the input with new template vars
  useEffect(() => {
    setDeployInput((prev) => {
      const newTemplateVars = {} as Record<string, string | number>
      templateVars.forEach((v) => {
        newTemplateVars[v.name] = v.value
      })
      return {
        ...prev,
        templateVars: newTemplateVars,
      }
    })
  }, [templateVars])

  useEffect(() => {
    // Isn't loading, isn't already loaded, isn't errored, and has id
    if (!blueprintLoading && !blueprintData && !blueprintError && id)
      getBlueprint({
        variables: {
          id,
        },
      }).catch(console.error)
  }, [id])

  useEffect(() => {
    if (deployBlueprintLoading)
      enqueueSnackbar({ message: 'Deploying blueprint...', variant: 'info' })
    if (deployBlueprintError)
      enqueueSnackbar({
        message: `Error deploying blueprint: ${deployBlueprintError.message}`,
        variant: 'error',
      })
  }, [deployBlueprintError, deployBlueprintLoading])

  useEffect(() => {
    if (deployBlueprintData) {
      navigate('/deployments')
      enqueueSnackbar({ message: 'Deployed blueprint!', variant: 'success' })
    }
  }, [deployBlueprintData])

  return (
    <Container
      sx={{
        display: 'flex',
        flexDirection: 'column',
        py: 2,
      }}
    >
      <Box sx={{ width: '100%', display: 'flex', alignItems: 'center' }}>
        <Typography variant="h4">
          Request "{blueprintData?.blueprint.name}"
        </Typography>
      </Box>
      <Divider sx={{ width: '100%', my: 2 }} />
      <Stack spacing={2}>
        <Typography variant="h6">Description</Typography>
        <Typography variant="body1">
          {blueprintData?.blueprint.description}
        </Typography>
        <Typography variant="h6">Quota Usage</Typography>
        <TableContainer component={Paper} elevation={0}>
          <TableHead>
            <TableRow>
              <TableCell>Resource</TableCell>
              <TableCell align="center">CPU Quota Requirements</TableCell>
              <TableCell align="center">RAM Quota Requirements (MiB)</TableCell>
              <TableCell align="center">
                Disk Quota Requirements (MiB)
              </TableCell>
              <TableCell align="center">Network Quota Requirements</TableCell>
              <TableCell align="center">Router Quota Requirements</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {blueprintData?.blueprint.resources.map((resource) => (
              <TableRow key={resource.id}>
                <TableCell sx={{ fontWeight: 'bold' }}>
                  {resource.key}
                </TableCell>
                <TableCell align="center">
                  {resource.quotaRequirements.cpu}
                </TableCell>
                <TableCell align="center">
                  {resource.quotaRequirements.ram}
                </TableCell>
                <TableCell align="center">
                  {resource.quotaRequirements.disk}
                </TableCell>
                <TableCell align="center">
                  {resource.quotaRequirements.network}
                </TableCell>
                <TableCell align="center">
                  {resource.quotaRequirements.router}
                </TableCell>
              </TableRow>
            ))}
            <TableRow
              sx={{
                '& .MuiTableCell-root': {
                  borderTop: '1px solid',
                  borderBottom: 0,
                },
              }}
            >
              <TableCell sx={{ fontWeight: 'bold' }}>Total</TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                {blueprintData?.blueprint.resources.reduce(
                  (a, b) => a + b.quotaRequirements.cpu,
                  0
                )}
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                {blueprintData?.blueprint.resources.reduce(
                  (a, b) => a + b.quotaRequirements.ram,
                  0
                )}
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                {blueprintData?.blueprint.resources.reduce(
                  (a, b) => a + b.quotaRequirements.disk,
                  0
                )}
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                {blueprintData?.blueprint.resources.reduce(
                  (a, b) => a + b.quotaRequirements.network,
                  0
                )}
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                {blueprintData?.blueprint.resources.reduce(
                  (a, b) => a + b.quotaRequirements.router,
                  0
                )}
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell sx={{ fontWeight: 'bold' }}>
                New Project Quota Usage
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                <Typography
                  fontWeight="bold"
                  color={
                    selectedProject &&
                    selectedProject.usageCpu +
                      (blueprintData?.blueprint.resources.reduce(
                        (a, b) => a + b.quotaRequirements.cpu,
                        0
                      ) ?? 0) >
                      selectedProject.quotaCpu
                      ? 'error'
                      : 'inherit'
                  }
                  component="span"
                  fontSize="inherit"
                >
                  {selectedProject
                    ? `${
                        selectedProject.usageCpu +
                        (blueprintData?.blueprint.resources.reduce(
                          (a, b) => a + b.quotaRequirements.cpu,
                          0
                        ) ?? 0)
                      } / ${selectedProject.quotaCpu}`
                    : '-'}
                </Typography>
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                <Typography
                  fontWeight="bold"
                  color={
                    selectedProject &&
                    selectedProject.usageRam +
                      (blueprintData?.blueprint.resources.reduce(
                        (a, b) => a + b.quotaRequirements.ram,
                        0
                      ) ?? 0) >
                      selectedProject.quotaRam
                      ? 'error'
                      : 'inherit'
                  }
                  component="span"
                  fontSize="inherit"
                >
                  {selectedProject
                    ? `${
                        selectedProject.usageRam +
                        (blueprintData?.blueprint.resources.reduce(
                          (a, b) => a + b.quotaRequirements.ram,
                          0
                        ) ?? 0)
                      } / ${selectedProject.quotaRam}`
                    : '-'}
                </Typography>
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                <Typography
                  fontWeight="bold"
                  color={
                    selectedProject &&
                    selectedProject.usageDisk +
                      (blueprintData?.blueprint.resources.reduce(
                        (a, b) => a + b.quotaRequirements.disk,
                        0
                      ) ?? 0) >
                      selectedProject.quotaDisk
                      ? 'error'
                      : 'inherit'
                  }
                  component="span"
                  fontSize="inherit"
                >
                  {selectedProject
                    ? `${
                        selectedProject.usageDisk +
                        (blueprintData?.blueprint.resources.reduce(
                          (a, b) => a + b.quotaRequirements.disk,
                          0
                        ) ?? 0)
                      } / ${selectedProject.quotaDisk}`
                    : '-'}
                </Typography>
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                <Typography
                  fontWeight="bold"
                  color={
                    selectedProject &&
                    selectedProject.usageNetwork +
                      (blueprintData?.blueprint.resources.reduce(
                        (a, b) => a + b.quotaRequirements.network,
                        0
                      ) ?? 0) >
                      selectedProject.quotaNetwork
                      ? 'error'
                      : 'inherit'
                  }
                  component="span"
                  fontSize="inherit"
                >
                  {selectedProject
                    ? `${
                        selectedProject.usageNetwork +
                        (blueprintData?.blueprint.resources.reduce(
                          (a, b) => a + b.quotaRequirements.network,
                          0
                        ) ?? 0)
                      } / ${selectedProject.quotaNetwork}`
                    : '-'}
                </Typography>
              </TableCell>
              <TableCell align="center" sx={{ fontWeight: 'bold' }}>
                <Typography
                  fontWeight="bold"
                  color={
                    selectedProject &&
                    selectedProject.usageRouter +
                      (blueprintData?.blueprint.resources.reduce(
                        (a, b) => a + b.quotaRequirements.router,
                        0
                      ) ?? 0) >
                      selectedProject.quotaRouter
                      ? 'error'
                      : 'inherit'
                  }
                  component="span"
                  fontSize="inherit"
                >
                  {selectedProject
                    ? `${
                        selectedProject.usageRouter +
                        (blueprintData?.blueprint.resources.reduce(
                          (a, b) => a + b.quotaRequirements.router,
                          0
                        ) ?? 0)
                      } / ${selectedProject.quotaRouter}`
                    : '-'}
                </Typography>
              </TableCell>
            </TableRow>
          </TableBody>
        </TableContainer>
        <Typography variant="h6">Project</Typography>
        <ProjectAutocomplete
          minRole="deployer"
          sx={{ flex: '1 1' }}
          onChange={(val) => {
            setDeployInput((prev) => ({ ...prev, projectId: val?.id || '' }))
            setSelectedProject(val)
          }}
          error={deployInput.projectId === ''}
          helperText={
            deployInput.projectId === '' ? 'A project is required' : undefined
          }
        />
        <Typography variant="h6">Variables</Typography>
        {templateVars.map((variable, i) => (
          <TextField
            key={`var_${variable.name}_${i}`}
            label={variable.name}
            variant="outlined"
            type={
              blueprintData?.blueprint.variableTypes[variable.name] === 'INT'
                ? 'number'
                : 'text'
            }
            value={variable.value}
            onChange={(e) => {
              setTemplateVars((prev) => {
                const newVars = [...prev]
                newVars[i].value = e.target.value
                return newVars
              })
            }}
            focused={variable.value === '' || variable.value === 0}
            color={
              variable.value === '' || variable.value === 0
                ? 'warning'
                : undefined
            }
            error={variable.value === '' || variable.value === 0}
          />
        ))}
      </Stack>
      <ContainerFab
        color="primary"
        variant="extended"
        sx={{ position: 'fixed', bottom: '2rem', right: '2rem' }}
        onClick={() => {
          if (id)
            deployBlueprint({
              variables: deployInput,
            }).catch(console.error)
        }}
        disabled={
          !deployInput.blueprintId ||
          !deployInput.projectId ||
          !deployInput.templateVars ||
          deployBlueprintLoading ||
          deployBlueprintData != null ||
          deployBlueprintError != null ||
          exceedsAnyQuota
        }
      >
        <Send sx={{ mr: 1 }} /> Request
      </ContainerFab>
    </Container>
  )
}
