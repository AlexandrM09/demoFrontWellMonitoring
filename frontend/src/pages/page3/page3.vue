<template>
<div class="title">  График
        <ChartWellDay> </ChartWellDay>
    <button >Randomize</button>
</div>
</template>
<script>
// eslint-disable-next-line spaced-comment
//import Chart from 'Chart.js/dist/Chart.js'
import ChartWellDay from '@/components/ChartWellDay'
export default {
  components: {
    ChartWellDay
  },
  data () {
    return {
      datacollection: {
        labels: [this.getRandomInt(), this.getRandomInt()],
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: [this.getRandomInt(), this.getRandomInt()]
          }, {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: [this.getRandomInt(), this.getRandomInt()]
          }
        ]
      }
    }
  },
  mounted () {
    // console.log(document.getElementById('bar-chart')) // I'm text inside the component.
  }
}
</script>
