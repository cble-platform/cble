import {
  Box,
  IconButton,
  Container,
  Divider,
  Grid,
  Typography,
  useMediaQuery,
  TextField,
  Autocomplete,
  LinearProgress,
  CircularProgress,
  Stack,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Button,
  Tooltip,
} from '@mui/material'
import { useCallback, useContext, useEffect, useMemo, useState } from 'react'
import {
  BlueprintInput,
  ListProvidersQuery,
  SearchProjectQuery,
  useCreateBlueprintMutation,
  useGetBlueprintLazyQuery,
  useListProvidersQuery,
  useSearchProjectLazyQuery,
  useUpdateBlueprintMutation,
} from '../../lib/api/generated'
import MonacoEditor, { useMonaco } from '@monaco-editor/react'
import { configureMonacoYaml } from 'monaco-yaml'
import { ThemeContext } from '../../theme'
import { Add, AutoFixHigh, Circle, Delete, Save } from '@mui/icons-material'
import { LoadingButton } from '@mui/lab'
import { useSnackbar } from 'notistack'
import { useNavigate, useParams } from 'react-router-dom'

export default function BlueprintForm({
  action,
}: {
  action: 'create' | 'edit'
}) {
  const { id } = useParams()
  const { themePreference } = useContext(ThemeContext)
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)')
  const { enqueueSnackbar } = useSnackbar()
  const navigate = useNavigate()
  const [blueprint, setBlueprint] = useState<BlueprintInput>({
    name: '',
    description: '',
    blueprintTemplate: `version: "1.0"\n`,
    providerId: '',
    variableTypes: {},
    projectId: '',
  })
  const [variableTypes, setVariableTypes] = useState<
    readonly {
      name: string
      type: string
    }[]
  >([])
  const monaco = useMonaco()
  const {
    data: providersData,
    error: providersError,
    loading: providersLoading,
  } = useListProvidersQuery()
  const [
    searchProjects,
    {
      data: searchProjectsData,
      error: searchProjectsError,
      loading: searchProjectsLoading,
    },
  ] = useSearchProjectLazyQuery()
  const [projectsSearchVal, setProjectsSearchVal] = useState<string>('')
  const [projectOptions, setProjectOptions] = useState<
    readonly SearchProjectQuery['searchProjects']['projects'][number][]
  >([])
  const [projectOpen, setProjectOpen] = useState<boolean>(false)
  // Create
  const [
    createBlueprint,
    {
      data: createBlueprintData,
      error: createBlueprintError,
      loading: createBlueprintLoading,
    },
  ] = useCreateBlueprintMutation()
  // Edit
  const [
    getBlueprint,
    {
      data: getBlueprintData,
      error: getBlueprintError,
      loading: getBlueprintLoading,
    },
  ] = useGetBlueprintLazyQuery()
  const [
    updateBlueprint,
    {
      data: updateBlueprintData,
      error: updateBlueprintError,
      loading: updateBlueprintLoading,
      reset: resetUpdateBlueprint,
    },
  ] = useUpdateBlueprintMutation()

  useEffect(() => {
    if (providersError)
      enqueueSnackbar({
        message: `Failed to load provider list: ${providersError.message}`,
        variant: 'error',
      })
    if (createBlueprintError)
      enqueueSnackbar({
        message: `Failed to create blueprint: ${createBlueprintError.message}`,
        variant: 'error',
      })
    if (getBlueprintError)
      enqueueSnackbar({
        message: `Failed to create blueprint: ${getBlueprintError.message}`,
        variant: 'error',
      })
    if (updateBlueprintError)
      enqueueSnackbar({
        message: `Failed to update blueprint: ${updateBlueprintError.message}`,
        variant: 'error',
      })
  }, [
    providersError,
    createBlueprintError,
    getBlueprintError,
    updateBlueprintError,
  ])

  // If we're editing the blueprint, fetch based on ID
  useEffect(() => {
    if (action === 'edit' && id)
      getBlueprint({
        variables: { id },
      }).catch(console.error)
  }, [id, action, getBlueprint])

  // Once we get edit data, update local form state
  useEffect(() => {
    if (getBlueprintData?.blueprint) {
      setBlueprint({
        name: getBlueprintData.blueprint.name,
        description: getBlueprintData.blueprint.description,
        blueprintTemplate: getBlueprintData.blueprint.blueprintTemplate,
        providerId: getBlueprintData.blueprint.provider.id,
        projectId: getBlueprintData.blueprint.project.id,
        variableTypes: getBlueprintData.blueprint.variableTypes,
      } as BlueprintInput)
      setVariableTypes(
        Object.keys(
          getBlueprintData.blueprint.variableTypes as Record<string, string>
        ).map((v) => ({
          name: v,
          type: getBlueprintData.blueprint.variableTypes[v],
        }))
      )
    }
  }, [getBlueprintData])

  // Move to edit page after creation
  useEffect(() => {
    if (createBlueprintData?.createBlueprint.id) {
      enqueueSnackbar({ message: 'Created blueprint!', variant: 'success' })
      navigate(`/blueprints/edit/${createBlueprintData.createBlueprint.id}`)
    }
  }, [createBlueprintData, navigate])

  // Show pop up when updating blueprint
  useEffect(() => {
    if (updateBlueprintData?.updateBlueprint.id) {
      enqueueSnackbar({ message: 'Updated blueprint!', variant: 'success' })
      resetUpdateBlueprint()
    }
  }, [updateBlueprintData, enqueueSnackbar, resetUpdateBlueprint])

  // Query for project autocomplete
  useEffect(() => {
    searchProjects({ variables: { search: projectsSearchVal } })
  }, [projectsSearchVal])

  // Set autocomplete values
  useEffect(() => {
    if (searchProjectsData?.searchProjects.projects)
      setProjectOptions(searchProjectsData.searchProjects.projects)
  }, [searchProjectsData])

  // Set variable types in input object
  useEffect(() => {
    const vars: Record<string, string> = {}
    variableTypes.forEach((v) => {
      vars[v.name] = v.type
    })
    setBlueprint((prev) => ({
      ...prev,
      variableTypes: vars,
    }))
  }, [variableTypes])

  const isVarDuplicate = useCallback(
    (index: number) => {
      return (
        variableTypes.filter((v) => v.name === variableTypes[index].name)
          .length > 1
      )
    },
    [variableTypes]
  )

  const isVarUsed = useCallback(
    (index: number) => {
      return new RegExp(`{{ ?.${variableTypes[index].name} ?}}`, 'g').test(
        blueprint.blueprintTemplate
      )
    },
    [blueprint, variableTypes]
  )

  const selectedProject = useMemo(() => {
    if (blueprint.projectId)
      return projectOptions.find((p) => p.id === blueprint.projectId)
    return null
  }, [projectOptions, blueprint])

  const handlePrettify = () => {
    if (monaco)
      monaco.editor.getEditors().forEach((editor) => {
        editor
          .getAction('editor.action.formatDocument')
          ?.run()
          .catch(console.error)
      })
  }

  const handleSubmitBlueprint = () => {
    if (action === 'create')
      createBlueprint({
        variables: {
          input: blueprint,
        },
      }).catch(console.error)
    else if (action === 'edit' && id)
      updateBlueprint({
        variables: {
          id,
          input: blueprint,
        },
      }).catch(console.error)
  }

  return (
    <Container
      sx={{ display: 'flex', flexDirection: 'column', height: '100%', py: 2 }}
    >
      <Stack
        direction="row"
        alignItems="center"
        justifyContent="space-between"
        spacing={4}
      >
        <TextField
          variant="standard"
          placeholder="Untitled Blueprint"
          value={blueprint.name}
          onChange={(e) =>
            setBlueprint((prev) => ({ ...prev, name: e.target.value }))
          }
          sx={{ flex: '1 1' }}
        ></TextField>
        <Autocomplete
          fullWidth
          disablePortal
          autoComplete
          clearOnEscape
          clearOnBlur={false}
          filterOptions={(x) => x}
          open={projectOpen}
          onOpen={() => {
            setProjectOpen(true)
          }}
          onClose={() => {
            setProjectOpen(false)
          }}
          loading={searchProjectsLoading}
          options={projectOptions}
          getOptionLabel={(option) => option.name}
          sx={{ flex: '1 1' }}
          isOptionEqualToValue={(option, val) =>
            val === undefined || option.id === val.id
          }
          value={selectedProject}
          onChange={(_, val) => {
            console.log(val)
            setBlueprint((prev) => ({ ...prev, projectId: val?.id ?? '' }))
          }}
          onInputChange={(_, val) => setProjectsSearchVal(val)}
          inputValue={projectsSearchVal}
          renderOption={(props, option) => (
            <li {...props} key={option.id}>
              <Typography>{option.name}</Typography>
            </li>
          )}
          renderInput={(params) => (
            <TextField
              {...params}
              label="Project"
              InputProps={{
                ...params.InputProps,
                endAdornment: (
                  <>
                    {searchProjectsLoading ? (
                      <CircularProgress color="inherit" size={20} />
                    ) : null}
                    {params.InputProps.endAdornment}
                  </>
                ),
              }}
            />
          )}
        />
        <Autocomplete
          disablePortal
          autoComplete
          clearOnEscape
          options={providersData?.providers.providers ?? []}
          getOptionKey={(
            option: ListProvidersQuery['providers']['providers'][number]
          ) => `${option.id}`}
          getOptionLabel={(
            option: ListProvidersQuery['providers']['providers'][number]
          ) => `${option.displayName} (${option.providerVersion})`}
          sx={{ flex: '1 1' }}
          value={
            (providersData?.providers.providers.find(
              (p) => p.id === blueprint.providerId
            ) ||
              null) ??
            null
          }
          onChange={(_, val) => {
            val && setBlueprint((prev) => ({ ...prev, providerId: val.id }))
          }}
          renderOption={(props, option) => (
            <li {...props}>
              <Circle
                color={option.isLoaded ? 'success' : 'error'}
                sx={{ height: '0.75rem' }}
              />
              {option.displayName} ({option.providerVersion})
            </li>
          )}
          renderInput={(params) => <TextField {...params} label="Provider" />}
          disabled={providersLoading || providersError !== undefined}
        />
        <Stack direction="row">
          <IconButton onClick={handlePrettify} sx={{ mr: 1 }} title="Prettify">
            <AutoFixHigh />
          </IconButton>
          <LoadingButton
            variant="contained"
            startIcon={action === 'create' ? <Add /> : <Save />}
            loading={createBlueprintLoading || updateBlueprintLoading}
            disabled={
              blueprint.name === '' ||
              blueprint.blueprintTemplate === '' ||
              blueprint.providerId === '' ||
              createBlueprintData != null
            }
            onClick={handleSubmitBlueprint}
          >
            {action === 'create' ? 'Create' : 'Save'}
          </LoadingButton>
        </Stack>
      </Stack>
      <Divider sx={{ my: 2 }} />
      {getBlueprintLoading && <LinearProgress />}
      <Grid container sx={{ flexGrow: 1 }} spacing={2}>
        <Grid
          item
          md={6}
          sx={{ display: 'flex', flexDirection: 'column', gap: 2 }}
        >
          <Typography variant="h6">Description</Typography>
          <TextField
            multiline
            sx={{ width: '100%' }}
            value={blueprint.description}
            onChange={(e) =>
              setBlueprint((prev) => ({ ...prev, description: e.target.value }))
            }
          />
          <Typography variant="h6">Variables</Typography>
          <Stack spacing={2}>
            {variableTypes.map((variable, i) => (
              <Stack
                direction="row"
                alignItems="center"
                spacing={2}
                key={`var_${i}`}
              >
                <FormControl variant="standard" sx={{ flex: '1 1' }}>
                  <Tooltip
                    followCursor
                    arrow
                    title={
                      isVarDuplicate(i)
                        ? `Variable ${variable.name} is already defined`
                        : !isVarUsed(i)
                          ? 'Variable is never used'
                          : ''
                    }
                  >
                    <TextField
                      label="Name"
                      variant="outlined"
                      value={variable.name}
                      onChange={(e) => {
                        setVariableTypes((prev) => {
                          const newVars = [...prev]
                          newVars[i].name = e.target.value
                          return newVars
                        })
                      }}
                      focused={!isVarUsed(i)}
                      color={!isVarUsed(i) ? 'warning' : 'primary'}
                      error={isVarDuplicate(i)}
                    />
                  </Tooltip>
                </FormControl>
                <FormControl variant="outlined" sx={{ flex: '1 1' }}>
                  <InputLabel id="Variable Type">Type</InputLabel>
                  <Select
                    labelId="demo-simple-select-standard-label"
                    id="demo-simple-select-standard"
                    value={variable.type}
                    onChange={(e) => {
                      setVariableTypes((prev) => {
                        const newVars = [...prev]
                        newVars[i].type = e.target.value
                        return newVars
                      })
                    }}
                    label="Type"
                  >
                    <MenuItem value="STRING">String</MenuItem>
                    <MenuItem value="INT">Integer</MenuItem>
                  </Select>
                </FormControl>
                <IconButton
                  color="error"
                  sx={{ marginTop: '1' }}
                  onClick={() =>
                    setVariableTypes((prev) => {
                      const newVars = [...prev]
                      newVars.splice(i, 1)
                      return newVars
                    })
                  }
                >
                  <Delete />
                </IconButton>
              </Stack>
            ))}
            <Button
              variant="contained"
              startIcon={<Add />}
              onClick={() =>
                setVariableTypes((prev) => [
                  ...prev,
                  { name: `var${prev.length + 1}`, type: 'STRING' },
                ])
              }
            >
              Add Variable
            </Button>
          </Stack>
        </Grid>
        <Grid item md={6}>
          <Typography variant="h6">Template</Typography>
          <Box
            sx={{
              maxHeight: 'calc(100dvh - 14rem)',
              height: '100%',
              borderWidth: 1,
              borderStyle: 'solid',
              borderColor: 'primary',
            }}
          >
            <MonacoEditor
              theme={
                themePreference === 'auto'
                  ? prefersDarkMode
                    ? 'vs-dark'
                    : 'light'
                  : themePreference === 'dark'
                    ? 'vs-dark'
                    : 'light'
              }
              language="yaml"
              options={{
                minimap: { enabled: false },
                scrollBeyondLastLine: false,
                wordWrap: 'on',
              }}
              value={blueprint.blueprintTemplate}
              onChange={(value, _) =>
                setBlueprint(
                  (prev) =>
                    ({
                      ...prev,
                      blueprintTemplate: value ?? '',
                    }) as BlueprintInput
                )
              }
              beforeMount={(monaco) => {
                configureMonacoYaml(monaco, {
                  enableSchemaRequest: true,
                  schemas: [
                    {
                      fileMatch: ['*'],
                      schema: {
                        type: 'object',
                        properties: {
                          version: {
                            type: 'string',
                            description: 'The blueprint syntax version',
                          },
                        },
                      },
                      uri: 'https://github.com/cble-platform',
                    },
                  ],
                })
              }}
            />
          </Box>
        </Grid>
      </Grid>
    </Container>
  )
}
