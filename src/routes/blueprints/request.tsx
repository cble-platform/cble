import {
  Box,
  Container,
  Divider,
  Stack,
  TextField,
  Typography,
} from '@mui/material'
import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import {
  useDeployBlueprintMutation,
  useGetBlueprintLazyQuery,
} from '../../lib/api/generated'
import { Send } from '@mui/icons-material'
import { useSnackbar } from 'notistack'
import ContainerFab from '../../components/container-fab'
import ProjectAutocomplete from '@/components/project-autocomplete'

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
      sx={{ display: 'flex', flexDirection: 'column', height: '100%', py: 2 }}
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
        <Typography variant="h6">Project</Typography>

        <ProjectAutocomplete
          minRole="deployer"
          sx={{ flex: '1 1' }}
          onChange={(val) =>
            setDeployInput((prev) => ({ ...prev, projectId: val || '' }))
          }
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
          deployBlueprintError != null
        }
      >
        <Send sx={{ mr: 1 }} /> Request
      </ContainerFab>
    </Container>
  )
}
