<template>
<div class="title"> у3
    <v-card
    class="mx-auto"
      max-width="344"
      outlined
      elevation=5
    >
        <ChartWellNPV> </ChartWellNPV>
    </v-card>
</div>
</template>
<script>
// eslint-disable-next-line spaced-comment
//import Chart from 'Chart.js/dist/Chart.js'
import ChartWellNPV from '@/components/ChartWellNPV'
export default {
  components: {
    ChartWellNPV
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
