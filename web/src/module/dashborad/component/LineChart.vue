<template>
  <div ref="chartRef" class="chart-container"></div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch } from 'vue'
import * as echarts from 'echarts'

const props = defineProps({
  title: {
    type: String,
    default: 'Line Chart',
  },
  data: {
    type: Array,
    default: () => [],
  },
  xAxisData: {
    type: Array,
    default: () => [],
  },
})

const chartRef = ref(null)
let chartInstance = null

const initChart = () => {
  if (chartRef.value) {
    chartInstance = echarts.init(chartRef.value)
    const option = {
      title: {
        text: props.title,
        left: 'center',
      },
      tooltip: {
        trigger: 'axis',
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true,
      },
      xAxis: {
        type: 'category',
        boundaryGap: false,
        data: props.xAxisData.length
          ? props.xAxisData
          : ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
      },
      yAxis: {
        type: 'value',
      },
      series: [
        {
          name: 'Data',
          type: 'line',
          smooth: true,
          data: props.data.length
            ? props.data
            : [150, 230, 224, 218, 135, 147, 260],
          areaStyle: {
            opacity: 0.1,
          },
          lineStyle: {
            width: 3,
          },
        },
      ],
    }
    chartInstance.setOption(option)
  }
}

const resizeChart = () => {
  chartInstance?.resize()
}

onMounted(() => {
  initChart()
  window.addEventListener('resize', resizeChart)
})

onBeforeUnmount(() => {
  window.removeEventListener('resize', resizeChart)
  chartInstance?.dispose()
})

watch(
  () => props.data,
  () => {
    if (chartInstance) {
      chartInstance.setOption({
        series: [
          {
            data: props.data,
          },
        ],
      })
    }
  },
)
</script>

<style scoped>
.chart-container {
  width: 100%;
  height: 400px;
}
</style>
