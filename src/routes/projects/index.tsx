import { useProjectsQuery } from '@/lib/api/generated'
import { Container, Box, Typography, Divider, Grid } from '@mui/material'
import { PieChart } from '@mui/x-charts'

// Format megabytes with units and auto-divide into larger units
function formatMegabytes(mib: number): string {
  if (mib < 1024) return `${mib} MiB`
  const gib = mib / 1024
  if (gib < 1024) return `${gib} GiB`
  const tib = mib / 1024
  if (tib < 1024) return `${tib} GiB`
  return `${mib}`
}

const pieChartPalette = ['#F76902', '#ddd']
const pieChartProps = {
  colors: pieChartPalette,
  height: 100,
  margin: { right: 5 },
  slotProps: {
    legend: { hidden: true },
  },
}

export default function Projects() {
  const {
    data: projectsData,
    error: projectsError,
    loading: projectsLoading,
  } = useProjectsQuery({ variables: { minRole: 'viewer' } })

  return (
    <Container
      sx={{
        display: 'flex',
        flexDirection: 'column',
        py: 2,
      }}
    >
      <Box sx={{ width: '100%', display: 'flex', alignItems: 'center' }}>
        <Typography variant="h4">Projects</Typography>
      </Box>
      <Divider sx={{ width: '100%', my: 2 }} />
      {projectsData?.projects.projects.map((project) => (
        <Grid container>
          <Grid item xs={12}>
            <Typography variant="h5">{project.name}</Typography>
          </Grid>
          <Grid item xs={12} sx={{ marginBottom: 2 }} container>
            <Grid
              xs
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 1,
              }}
            >
              <PieChart
                series={[
                  {
                    paddingAngle: 0,
                    data: [
                      {
                        label: 'Used',
                        value: project.usageCpu,
                      },
                      {
                        label: 'Unused',
                        value: project.quotaCpu,
                      },
                    ],
                  },
                ]}
                {...pieChartProps}
              />
              <Typography variant="subtitle1">CPU Quota</Typography>
              <Typography variant="subtitle1">
                {project.usageCpu} of {project.quotaCpu}
              </Typography>
            </Grid>
            <Grid
              xs
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 1,
              }}
            >
              <PieChart
                series={[
                  {
                    data: [
                      {
                        label: 'Used',
                        value: project.usageRam,
                      },
                      {
                        label: 'Unused',
                        value: project.quotaRam,
                      },
                    ],
                  },
                ]}
                {...pieChartProps}
              />
              <Typography variant="subtitle1">RAM Quota</Typography>
              <Typography variant="subtitle1">
                {formatMegabytes(project.usageRam)} of{' '}
                {formatMegabytes(project.quotaRam)}
              </Typography>
            </Grid>
            <Grid
              xs
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 1,
              }}
            >
              <PieChart
                series={[
                  {
                    data: [
                      {
                        label: 'Used',
                        value: project.usageDisk,
                      },
                      {
                        label: 'Unused',
                        value: project.quotaDisk,
                      },
                    ],
                  },
                ]}
                {...pieChartProps}
              />
              <Typography variant="subtitle1">Disk Quota</Typography>
              <Typography variant="subtitle1">
                {formatMegabytes(project.usageDisk)} of{' '}
                {formatMegabytes(project.quotaDisk)}
              </Typography>
            </Grid>
            <Grid
              xs
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 1,
              }}
            >
              <PieChart
                series={[
                  {
                    data: [
                      {
                        label: 'Used',
                        value: project.usageNetwork,
                      },
                      {
                        label: 'Unused',
                        value: project.quotaNetwork,
                      },
                    ],
                  },
                ]}
                {...pieChartProps}
              />
              <Typography variant="subtitle1">Network Quota</Typography>
              <Typography variant="subtitle1">
                {project.usageNetwork} of {project.quotaNetwork}
              </Typography>
            </Grid>
            <Grid
              xs
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 1,
              }}
            >
              <PieChart
                series={[
                  {
                    data: [
                      {
                        label: 'Used',
                        value: project.usageRouter,
                      },
                      {
                        label: 'Unused',
                        value: project.quotaRouter,
                      },
                    ],
                  },
                ]}
                {...pieChartProps}
              />
              <Typography variant="subtitle1">Router Quota</Typography>
              <Typography variant="subtitle1">
                {project.usageRouter} of {project.quotaRouter}
              </Typography>
            </Grid>
          </Grid>
        </Grid>
      ))}
    </Container>
  )
}
