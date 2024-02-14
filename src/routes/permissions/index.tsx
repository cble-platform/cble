import {
  Action,
  ListPermissionsQuery,
  ObjectType,
  SearchGroupsQuery,
  SearchUsersQuery,
  SubjectType,
  useGrantPermissionMutation,
  useListPermissionsQuery,
  useMeHasPermissionQuery,
  useRevokePermissionMutation,
  useSearchGroupsLazyQuery,
  useSearchUsersLazyQuery,
} from '@/lib/api/generated'
import { Add, Cancel, Group, Person } from '@mui/icons-material'
import {
  Autocomplete,
  Box,
  Button,
  ButtonGroup,
  CircularProgress,
  Container,
  Divider,
  FormControl,
  InputLabel,
  MenuItem,
  Modal,
  OutlinedInput,
  Paper,
  Select,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableFooter,
  TableHead,
  TablePagination,
  TableRow,
  TextField,
  Typography,
} from '@mui/material'
import { useEffect, useMemo, useState } from 'react'
import { modalDefault } from '@/lib/modal'
import { UUIDTextMask } from '@/lib/input'
import { useSnackbar } from 'notistack'

export default function Permissions() {
  const {
    data: listPermissionsData,
    error: listPermissionsError,
    loading: listPermissionsLoading,
    refetch: refetchListPermissions,
  } = useListPermissionsQuery()
  const { data: hasGrantPermissionData } = useMeHasPermissionQuery({
    variables: {
      objectType: ObjectType.Permission,
      objectID: null,
      action: Action.PermissionGrant,
    },
  })
  const [grantModalOpen, setGrantModalOpen] = useState<boolean>(false)
  const [revokeModalOpen, setRevokeModalOpen] = useState<boolean>(false)
  const [revokeModalData, setRevokeModalData] =
    useState<ListPermissionsQuery['permissions']['permissions'][number]>()
  const [page, setPage] = useState<number>(0)
  const [rowsPerPage, setRowsPerPage] = useState<number>(10)

  const emptyRows = useMemo(
    () =>
      page > 0
        ? Math.max(
            0,
            rowsPerPage -
              (listPermissionsData?.permissions.permissions.length ?? 0)
          )
        : 0,
    [page, rowsPerPage, listPermissionsData]
  )

  useEffect(() => {
    refetchListPermissions({
      count: rowsPerPage,
      offset: rowsPerPage * page,
    })
  }, [page, rowsPerPage])

  return (
    <Container sx={{ py: 3 }}>
      <Box
        sx={{
          display: 'flex',
          alignContent: 'center',
          justifyContent: 'space-between',
        }}
      >
        <Typography variant="h4">Permissions</Typography>
        {hasGrantPermissionData?.meHasPermission && (
          <Button
            onClick={() => setGrantModalOpen(true)}
            variant="contained"
            color="primary"
            startIcon={<Add />}
          >
            Grant Permission
          </Button>
        )}
      </Box>
      <Divider sx={{ my: 3 }} />
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
          <TableHead>
            <TableRow>
              <TableCell>Permission</TableCell>
              <TableCell />
            </TableRow>
          </TableHead>
          <TableBody>
            {listPermissionsData?.permissions.permissions.map((row) => (
              <TableRow
                key={row.id}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.displayString}
                </TableCell>
                <TableCell align="right">
                  <Button
                    size="small"
                    startIcon={<Cancel />}
                    color="error"
                    onClick={() => {
                      setRevokeModalData(row)
                      setRevokeModalOpen(true)
                    }}
                  >
                    Revoke
                  </Button>
                </TableCell>
              </TableRow>
            ))}
            {emptyRows > 0 && (
              <TableRow
                style={{
                  height: 47 * emptyRows, // 55 for not dense
                }}
              >
                <TableCell colSpan={2} />
              </TableRow>
            )}
          </TableBody>
          <TableFooter>
            <TableRow>
              <TablePagination
                rowsPerPageOptions={[10, 25, 50]}
                rowsPerPage={rowsPerPage}
                onRowsPerPageChange={(e) => {
                  setPage(0)
                  setRowsPerPage(parseInt(e.target.value))
                }}
                page={page}
                onPageChange={(_e, value) => setPage(value)}
                count={listPermissionsData?.permissions.total ?? 0}
                disabled={listPermissionsLoading}
              />
            </TableRow>
          </TableFooter>
        </Table>
      </TableContainer>
      <Modal open={grantModalOpen} onClose={() => setGrantModalOpen(false)}>
        <Box sx={{ ...modalDefault }}>
          {grantModalOpen && (
            <GrantModal
              onClose={() => {
                setGrantModalOpen(false)
                setPage(0)
                refetchListPermissions()
              }}
            />
          )}
        </Box>
      </Modal>
      <Modal open={revokeModalOpen} onClose={() => setRevokeModalOpen(false)}>
        <Box sx={{ ...modalDefault }}>
          {revokeModalOpen && (
            <RevokeModal
              onClose={() => {
                setRevokeModalOpen(false)
                setPage(0)
                refetchListPermissions()
              }}
              data={revokeModalData}
            />
          )}
        </Box>
      </Modal>
    </Container>
  )
}

type SearchUserResult = SearchUsersQuery['searchUsers']['users'][number]
type SearchGroupResult = SearchGroupsQuery['searchGroups']['groups'][number]

const isUserResult = (
  r: SearchUserResult | SearchGroupResult
): r is SearchUserResult => (r as SearchUserResult).username !== undefined

function GrantModal({ onClose }: { onClose: () => void }): React.ReactElement {
  const { enqueueSnackbar } = useSnackbar()
  const [
    searchUsers,
    {
      data: searchUsersData,
      error: searchUsersError,
      loading: searchUsersLoading,
    },
  ] = useSearchUsersLazyQuery()
  const [
    searchGroups,
    {
      data: searchGroupsData,
      error: searchGroupsError,
      loading: searchGroupsLoading,
    },
  ] = useSearchGroupsLazyQuery()
  const [subjectOptions, setSubjectOptions] = useState<
    readonly SearchUserResult[] | readonly SearchGroupResult[]
  >([])
  const [subjectOpen, setSubjectOpen] = useState<boolean>(false)
  const [subjectSearchVal, setSubjectSearchVal] = useState<string>('')
  const [subjectType, setSubjectType] = useState<SubjectType>(SubjectType.User)
  const [subjectID, setSubjectID] = useState<string>('')
  const [objectType, setObjectType] = useState<ObjectType>(ObjectType.Blueprint)
  const [objectID, setObjectID] = useState<string>('')
  const [action, setAction] = useState<Action>(Action.BlueprintCreate)
  const [
    grantPermission,
    {
      data: grantPermissionData,
      error: grantPermissionError,
      loading: grantPermissionLoading,
    },
  ] = useGrantPermissionMutation()

  const selectedSubject = useMemo(() => {
    if (subjectType === SubjectType.User && searchUsersData && subjectID)
      return searchUsersData.searchUsers.users.find((u) => u.id === subjectID)
    if (subjectType === SubjectType.Group && searchGroupsData && subjectID)
      return searchGroupsData.searchGroups.groups.find(
        (g) => g.id === subjectID
      )
    return null
  }, [subjectType, subjectID])

  useEffect(() => {
    if (subjectType === SubjectType.User && searchUsersData?.searchUsers.users)
      setSubjectOptions(searchUsersData.searchUsers.users)
    else if (
      subjectType === SubjectType.Group &&
      searchGroupsData?.searchGroups.groups
    )
      setSubjectOptions(searchGroupsData.searchGroups.groups)
  }, [searchUsersData, searchGroupsData])

  useEffect(() => {
    if (subjectType === SubjectType.User)
      searchUsers({ variables: { search: subjectSearchVal } })
    else searchGroups({ variables: { search: subjectSearchVal } })
  }, [subjectSearchVal])

  useEffect(() => {
    if (grantPermissionError)
      enqueueSnackbar({
        message: grantPermissionError.message,
        variant: 'error',
      })
  }, [grantPermissionError])

  useEffect(() => {
    if (grantPermissionData?.grantPermission?.id) {
      enqueueSnackbar({
        message: `Granted permission ${grantPermissionData.grantPermission.displayString}`,
        variant: 'success',
      })
      onClose()
    }
  }, [grantPermissionData])

  return (
    <Stack gap={2}>
      <Stack direction="row" gap={2}>
        <ButtonGroup variant="outlined" aria-label="Basic button group">
          <Button
            startIcon={<Person />}
            variant={
              subjectType === SubjectType.User ? 'contained' : 'outlined'
            }
            onClick={() => {
              setSubjectType(SubjectType.User)
              setSubjectID('')
              setSubjectOptions([])
            }}
          >
            User
          </Button>
          <Button
            startIcon={<Group />}
            variant={
              subjectType === SubjectType.Group ? 'contained' : 'outlined'
            }
            onClick={() => {
              setSubjectType(SubjectType.Group)
              setSubjectID('')
              setSubjectOptions([])
            }}
          >
            Group
          </Button>
        </ButtonGroup>
        <Autocomplete
          fullWidth
          disablePortal
          autoComplete
          clearOnEscape
          filterOptions={(x) => x}
          open={subjectOpen}
          onOpen={() => {
            setSubjectOpen(true)
          }}
          onClose={() => {
            setSubjectOpen(false)
          }}
          loading={searchUsersLoading || searchGroupsLoading}
          options={subjectOptions}
          getOptionLabel={(option) =>
            isUserResult(option)
              ? `${option.firstName} ${option.lastName} (${option.email})`
              : `${option.name}`
          }
          // sx={{ width: 300 }}
          isOptionEqualToValue={(option, val) =>
            val === undefined || option.id === val.id
          }
          value={selectedSubject}
          onChange={(_, val) => setSubjectID(val?.id ?? '')}
          onInputChange={(_, val) => setSubjectSearchVal(val)}
          inputValue={subjectSearchVal}
          renderOption={(props, option) =>
            isUserResult(option) ? (
              <li {...props} key={option.id}>
                <Typography>
                  {option.firstName} {option.lastName}
                </Typography>
                {/* {option.displayName} ({option.providerVersion}) */}
              </li>
            ) : (
              <li {...props}>{option.name}</li>
            )
          }
          renderInput={(params) => (
            <TextField
              {...params}
              label={subjectType === SubjectType.User ? 'User' : 'Group'}
              InputProps={{
                ...params.InputProps,
                endAdornment: (
                  <>
                    {searchUsersLoading || searchGroupsLoading ? (
                      <CircularProgress color="inherit" size={20} />
                    ) : null}
                    {params.InputProps.endAdornment}
                  </>
                ),
              }}
            />
          )}
          // disabled={providersLoading || providersError !== undefined}
        />
      </Stack>
      <Stack direction="row" gap={2}>
        <FormControl sx={{ minWidth: '10rem' }}>
          <InputLabel id="object-type-label">Object Type</InputLabel>
          <Select
            labelId="object-type-label"
            id="object-type"
            value={objectType}
            label="Object Type"
            onChange={(e) => {
              setObjectType(e.target.value as ObjectType)
              setAction(
                Object.values(Action).filter(
                  (a) => a.indexOf(e.target.value) === 0
                )[0] || Action.BlueprintCreate
              )
            }}
          >
            <MenuItem value={ObjectType.Blueprint}>Blueprint</MenuItem>
            <MenuItem value={ObjectType.Deployment}>Deployment</MenuItem>
            <MenuItem value={ObjectType.Group}>Group</MenuItem>
            <MenuItem value={ObjectType.Permission}>Permission</MenuItem>
            <MenuItem value={ObjectType.Provider}>Provider</MenuItem>
            <MenuItem value={ObjectType.User}>User</MenuItem>
          </Select>
        </FormControl>
        <FormControl variant="outlined" fullWidth>
          <InputLabel htmlFor="object-id">Object ID</InputLabel>
          <OutlinedInput
            value={objectID}
            onChange={(e) => setObjectID(e.target.value)}
            // name="textmask"
            label="Object ID"
            id="object-id"
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            inputComponent={UUIDTextMask as any}
          />
        </FormControl>
        {/* <TextField fullWidth label="Object ID" variant="outlined" /> */}
      </Stack>
      <Stack direction="row" gap={2}>
        <FormControl fullWidth sx={{ minWidth: '10rem' }}>
          <InputLabel id="object-type-label">Action</InputLabel>
          <Select
            labelId="object-type-label"
            id="object-type"
            value={action}
            label="Action"
            onChange={(e) => {
              setAction(e.target.value as Action)
            }}
          >
            {Object.values(Action)
              .filter((a) => a.indexOf(objectType) === 0)
              .map((a) => (
                <MenuItem key={a} value={a}>
                  {a}
                </MenuItem>
              ))}
          </Select>
        </FormControl>
      </Stack>
      <Stack direction="row" gap={2} justifyContent="flex-end">
        <Button
          color="secondary"
          onClick={onClose}
          disabled={grantPermissionLoading}
        >
          Cancel
        </Button>
        <Button
          variant="contained"
          disabled={grantPermissionLoading}
          onClick={() =>
            grantPermission({
              variables: {
                subjectType,
                subjectID,
                objectType,
                objectID: objectID === '*' ? null : objectID,
                action,
              },
            })
          }
        >
          Grant
        </Button>
      </Stack>
    </Stack>
  )
}

function RevokeModal({
  onClose,
  data,
}: {
  onClose: () => void
  data?: ListPermissionsQuery['permissions']['permissions'][number]
}): React.ReactElement {
  const { enqueueSnackbar } = useSnackbar()
  const [
    revokePermission,
    {
      data: revokePermissionData,
      error: revokePermissionError,
      loading: revokePermissionLoading,
    },
  ] = useRevokePermissionMutation()

  useEffect(() => {
    if (revokePermissionData?.revokePermission) {
      enqueueSnackbar({ message: `Revoked permission ${data?.displayString}` })
      onClose()
    }
  }, [revokePermissionData])

  return (
    <Stack gap={2}>
      {data && (
        <>
          <Typography>
            Are you sure you want to revoke {data.displayString}?
          </Typography>
          <Stack direction="row" justifyContent="flex-end" gap={2}>
            <Button
              color="secondary"
              onClick={onClose}
              disabled={revokePermissionLoading}
            >
              Cancel
            </Button>
            <Button
              variant="contained"
              disabled={revokePermissionLoading}
              color="error"
              onClick={() =>
                data &&
                revokePermission({
                  variables: {
                    subjectType: data.subjectType,
                    subjectID: data.subjectId,
                    objectType: data.objectType,
                    objectID: data.objectId,
                    action: data.action,
                  },
                })
              }
            >
              Revoke
            </Button>
          </Stack>
        </>
      )}
    </Stack>
  )
}
