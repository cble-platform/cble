import {
  SearchProjectQuery,
  useSearchProjectLazyQuery,
} from '@/lib/api/generated'
import {
  Autocomplete,
  Checkbox,
  CircularProgress,
  TextField,
  Typography,
} from '@mui/material'
import { useEffect, useState } from 'react'
import { SxProps, Theme } from '@mui/material/styles'

interface CommonProjectAutocompleteProps {
  sx?: SxProps<Theme>
  label?: string
  error?: boolean
  helperText?: string
  minRole?: 'admin' | 'developer' | 'deployer' | 'viewer'
}

interface MultipleProjectAutocompleteProps
  extends CommonProjectAutocompleteProps {
  multiple: true
  onChange: (projectIds: string[] | null) => void
}

interface SingleProjectAutocompleteProps
  extends CommonProjectAutocompleteProps {
  multiple?: false
  onChange: (projectId: string | null) => void
}

type ProjectAutocompleteProps =
  | SingleProjectAutocompleteProps
  | MultipleProjectAutocompleteProps

type OptionType = SearchProjectQuery['searchProjects']['projects'][number]

export default function ProjectAutocomplete({
  minRole,
  multiple,
  onChange,
  sx,
  label,
  error,
  helperText,
}: ProjectAutocompleteProps) {
  const [projectsSearchVal, setProjectsSearchVal] = useState<string>('')
  const [projectOpen, setProjectOpen] = useState<boolean>(false)
  const [projectOptions, setProjectOptions] = useState<readonly OptionType[]>(
    []
  )
  const [
    searchProjects,
    {
      data: searchProjectsData,
      error: searchProjectsError,
      loading: searchProjectsLoading,
    },
  ] = useSearchProjectLazyQuery()

  // Query for project autocomplete
  useEffect(() => {
    searchProjects({ variables: { search: projectsSearchVal, minRole } })
  }, [projectsSearchVal])

  // Set autocomplete values
  useEffect(() => {
    if (searchProjectsData?.searchProjects.projects)
      setProjectOptions(searchProjectsData.searchProjects.projects)
  }, [searchProjectsData])

  const handleChange = (val: OptionType | OptionType[] | null) => {
    if (multiple) onChange((val as OptionType[])?.map((v) => v.id))
    else onChange((val as OptionType)?.id)
  }

  return (
    <Autocomplete
      sx={sx}
      multiple={multiple}
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
      isOptionEqualToValue={(option, val) => option.id === val.id}
      onChange={(_, val) => handleChange(val)}
      onInputChange={(_, val) => setProjectsSearchVal(val)}
      inputValue={projectsSearchVal}
      renderOption={(props, option, { selected }) => (
        <li {...props} key={option.id}>
          {multiple && <Checkbox sx={{ mr: 1 }} checked={selected} />}
          <Typography>{option.name}</Typography>
        </li>
      )}
      renderInput={(params) => (
        <TextField
          {...params}
          label={label || 'Project'}
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
          error={error}
          helperText={helperText}
        />
      )}
    />
  )
}
