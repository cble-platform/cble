import { useSnackbar } from 'notistack'
import React, { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { GetDeploymentQuery, useGetDeploymentQuery } from '@/api/generated'
import {
  Container,
  Typography,
  Divider,
  LinearProgress,
  Box,
  Menu,
  MenuItem,
  Button,
  Card,
  CardContent,
} from '@mui/material'
import { ChevronLeft, ExpandMore } from '@mui/icons-material'
import ReactFlow, {
  Background,
  BackgroundVariant,
  ConnectionLineType,
  Controls,
  Edge,
  Handle,
  Node,
  Position,
  useEdgesState,
  useNodesState,
  useReactFlow,
} from 'reactflow'
import Dagre from '@dagrejs/dagre'

import 'reactflow/dist/style.css'

const dagreGraph = new Dagre.graphlib.Graph()
dagreGraph.setDefaultEdgeLabel(() => ({}))

const nodeWidth = 172
const nodeHeight = 36

const getLayoutedElements = (
  nodes: Node<GetDeploymentQuery['deployment']['deploymentNodes'][number]>[],
  edges: Edge[],
  options: { direction: 'TB' | 'BT' | 'LR' | 'RL' }
) => {
  const isHorizontal = options.direction === 'LR' || options.direction === 'RL'
  dagreGraph.setGraph({
    rankdir: options.direction,
    nodesep: nodeWidth / 2,
    edgesep: 25,
  })

  nodes.forEach((node) =>
    dagreGraph.setNode(node.id, {
      ...node,
      width: nodeWidth,
      height: nodeHeight,
    })
  )
  edges.forEach((edge) => dagreGraph.setEdge(edge.source, edge.target))

  Dagre.layout(dagreGraph)

  nodes.forEach((node) => {
    const nodeWithPosition = dagreGraph.node(node.id)
    node.targetPosition = isHorizontal ? Position.Left : Position.Top
    node.sourcePosition = isHorizontal ? Position.Right : Position.Bottom

    // We are shifting the dagre node position (anchor=center center) to the top left
    // so it matches the React Flow node anchor point (top left).
    node.position = {
      x: nodeWithPosition.x - nodeWidth / 2,
      y: nodeWithPosition.y - nodeHeight / 2,
    }

    return node
  })

  return { nodes, edges }
}

function generateFlowData(
  deploymentNodes: GetDeploymentQuery['deployment']['deploymentNodes']
): {
  nodes: Node<GetDeploymentQuery['deployment']['deploymentNodes'][number]>[]
  edges: Edge[]
} {
  const nodes: Node<
    GetDeploymentQuery['deployment']['deploymentNodes'][number]
  >[] = []
  const edges: Edge[] = []

  for (let i = 0; i < deploymentNodes.length; i++) {
    const node = deploymentNodes[i]
    nodes.push({
      id: node.id,
      position: { x: 100 * i, y: 0 },
      type: 'deploymentNode',
      data: node,
    })
    for (const nextNode of node.nextNodes) {
      edges.push({
        id: `${node.id}-${nextNode.id}`,
        source: node.id,
        target: nextNode.id,
        // type: 'smoothstep',
      })
    }
  }

  return {
    nodes,
    edges,
  }
}

const nodeTypes = { deploymentNode: DeploymentNodeNode }

function DeploymentNodeNode({
  data,
}: {
  data: GetDeploymentQuery['deployment']['deploymentNodes'][number]
}) {
  return (
    <Card variant="outlined" sx={{ width: nodeWidth, height: nodeHeight }}>
      <CardContent sx={{ p: '0.25rem !important' }}>
        <Handle type="target" position={Position.Top} />
        <Typography component="div" variant="h5" sx={{ fontSize: '0.75rem' }}>
          {data.resource.key}
        </Typography>
        <Handle type="source" position={Position.Bottom} />
      </CardContent>
    </Card>
  )
}

export default function DeploymentDetails() {
  const { id } = useParams()
  const { enqueueSnackbar } = useSnackbar()
  const { fitView } = useReactFlow()
  const navigate = useNavigate()
  const {
    data: getDeploymentData,
    error: getDeploymentError,
    loading: getDeploymentLoading,
  } = useGetDeploymentQuery({ variables: { id: id || '' } })
  const [moreMenuEl, setMoreMenuEl] = useState<null | HTMLElement>(null)
  const [nodes, setNodes, onNodesChange] = useNodesState<
    GetDeploymentQuery['deployment']['deploymentNodes'][number]
  >([])
  const [edges, setEdges, onEdgesChange] = useEdgesState([])

  useEffect(() => {
    if (getDeploymentError)
      enqueueSnackbar({
        message: `Failed to get deployment: ${getDeploymentError.message}`,
        variant: 'error',
      })
  }, [getDeploymentError])

  useEffect(() => {
    if (getDeploymentData) {
      const { nodes, edges } = generateFlowData(
        getDeploymentData.deployment.deploymentNodes
      )
      const layouted = getLayoutedElements(nodes, edges, { direction: 'TB' })
      setNodes([...layouted.nodes])
      setEdges([...layouted.edges])

      window.requestAnimationFrame(() => {
        fitView()
      })
    }
  }, [getDeploymentData])

  // const onLayout = useCallback(
  //   (direction: 'TB' | 'BT' | 'LR' | 'RL') => {
  //     const layouted = getLayoutedElements(nodes, edges, { direction })

  //     setNodes([...layouted.nodes])
  //     setEdges([...layouted.edges])

  //     window.requestAnimationFrame(() => {
  //       fitView()
  //     })
  //   },
  //   [nodes, edges]
  // )

  return (
    <Container sx={{ py: 3 }}>
      <Button href="/deployments" startIcon={<ChevronLeft />} sx={{ mb: 2 }}>
        Back
      </Button>
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
        }}
      >
        <Typography variant="h4">
          Details - {getDeploymentData?.deployment.name}
        </Typography>
        <Button
          id="more-button"
          aria-controls={moreMenuEl ? 'more-menu' : undefined}
          aria-haspopup="true"
          aria-expanded={moreMenuEl ? 'true' : undefined}
          onClick={(e) => setMoreMenuEl(e.currentTarget)}
          startIcon={<ExpandMore />}
        >
          Actions
        </Button>
        <Menu
          id="more-menu"
          anchorEl={moreMenuEl}
          open={Boolean(moreMenuEl)}
          onClose={() => setMoreMenuEl(null)}
          MenuListProps={{
            'aria-labelledby': 'more-button',
          }}
        >
          <MenuItem
            onClick={() =>
              navigate(
                `/deployments/destroy/${getDeploymentData?.deployment.id}`
              )
            }
          >
            Destroy
          </MenuItem>
        </Menu>
      </Box>
      <Divider sx={{ my: 2 }} />
      {getDeploymentLoading && <LinearProgress sx={{ my: 2 }} />}
      <Box
        sx={{
          display: 'grid',
          gridTemplateColumns: '1fr 3fr',
          gridTemplateRows: '1fr',
          gap: 2,
        }}
      >
        <Box></Box>
        <Box
          sx={{ width: '100%', height: '100%', minHeight: 500, px: 3, py: 1 }}
        >
          <ReactFlow
            nodes={nodes}
            edges={edges}
            fitView
            attributionPosition="top-right"
            nodeTypes={nodeTypes}
            connectionLineType={ConnectionLineType.SmoothStep}
          >
            <Controls />
            {/* <MiniMap /> */}
            <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
          </ReactFlow>
          {/* <MuiMarkdown>
            {getDeploymentData?.deployment.blueprint.description}
          </MuiMarkdown> */}
        </Box>
      </Box>
    </Container>
  )
}
