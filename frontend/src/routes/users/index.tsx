import {
  Box,
  Button,
  ButtonGroup,
  Container,
  Divider,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableFooter,
  TableHead,
  TablePagination,
  TableRow,
  Typography,
} from '@mui/material'
import { useSnackbar } from 'notistack'
import {
  Action,
  ListProvidersQuery,
  ObjectType,
  useListUsersQuery,
  useMeHasPermissionQuery,
} from '../../lib/api/generated'
import { useEffect, useMemo, useState } from 'react'
import { Add, Delete, Edit } from '@mui/icons-material'
import { useNavigate } from 'react-router-dom'

export default function Users() {
  const { enqueueSnackbar } = useSnackbar()
  const navigate = useNavigate()
  const { data: createUsersData } = useMeHasPermissionQuery({
    variables: {
      objectID: null,
      objectType: ObjectType.User,
      action: Action.UserCreate,
    },
  })
  const {
    data: listUsersData,
    error: listUsersError,
    loading: listUsersLoading,
    refetch: refetchListUsers,
  } = useListUsersQuery()
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null)
  const [moreMenuProvider, setMoreMenuProvider] =
    useState<ListProvidersQuery['providers']['providers'][number]>()
  const [page, setPage] = useState<number>(0)
  const [rowsPerPage, setRowsPerPage] = useState<number>(10)

  const emptyRows = useMemo(
    () =>
      page > 0
        ? Math.max(0, rowsPerPage - (listUsersData?.users.users.length ?? 0))
        : 0,
    [page, rowsPerPage, listUsersData]
  )

  useEffect(() => {
    refetchListUsers({
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
        <Typography variant="h4">Users</Typography>
        {createUsersData?.meHasPermission && (
          <Button
            href="/users/create"
            variant="contained"
            color="primary"
            startIcon={<Add />}
          >
            Create
          </Button>
        )}
      </Box>
      <Divider sx={{ my: 3 }} />

      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} size="small" aria-label="a dense table">
          <TableHead>
            <TableRow>
              <TableCell>Username</TableCell>
              <TableCell align="center">First Name</TableCell>
              <TableCell align="center">Last Name</TableCell>
              <TableCell align="center">Email</TableCell>
              <TableCell />
            </TableRow>
          </TableHead>
          <TableBody>
            {listUsersData?.users.users.map((row) => (
              <TableRow
                key={row.id}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.username}
                </TableCell>
                <TableCell align="center">{row.firstName}</TableCell>
                <TableCell align="center">{row.lastName}</TableCell>
                <TableCell align="center">{row.email}</TableCell>
                <TableCell align="right">
                  <ButtonGroup variant="text" aria-label="Controls">
                    <Button
                      size="small"
                      startIcon={<Edit />}
                      color="warning"
                      href={`/users/edit/${row.id}`}
                    >
                      Edit
                    </Button>
                    <Button
                      size="small"
                      startIcon={<Delete />}
                      color="error"
                      onClick={() => {
                        // setRevokeModalData(row)
                        // setRevokeModalOpen(true)
                      }}
                    >
                      Delete
                    </Button>
                  </ButtonGroup>
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
                count={listUsersData?.users.total ?? 0}
                disabled={listUsersLoading}
              />
            </TableRow>
          </TableFooter>
        </Table>
      </TableContainer>
    </Container>
  )
}
