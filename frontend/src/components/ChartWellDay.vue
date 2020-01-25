<script>
// import { Line } from 'vue-chartjs'
import {Bar, Line} from 'vue-chartjs'
// const { reactiveProp } = mixins

export default {
  extends: Bar,
  mixins: Line,
  props: ['options'],
  mounted () {
    this.renderChart({
      type: 'bar',
      labels: ['3.12', '4.12', '5.12', '6.12', '7.12', '8.12', '9.12',
        '10.12', '11.12', '12.12', '13.12', '14.12', '15.12', '16.12', '17.12', '18.12',
        '19.12', '20.12'],
      datasets: [
        {
          label: 'Фактическая',
          type: 'line',
          borderColor: '#00E676',
          data: [{y: 0}, {y: 120}, {y: 240}, {y: 400}, {y: 400}, {y: 400}, {y: 540},
            {y: 890}, {y: 1200}, {y: 1560}, {y: 2200}, {y: 2300}, {y: 2300}, {y: 2300}],
          fill: false
        },
        {
          label: 'Плановая',
          type: 'line',
          borderColor: '#AFB42B',
          data: [{x: '3.12', y: 0}, {x: '6.12', y: 400}, {x: '14.12', y: 2300}, {x: '20.12', y: 3100}],
          fill: false
        },
        {
          label: 'План:OK',
          type: 'bar',
          backgroundColor: '#3e95cd',
          data: [0, 0, 0, 400, 0, 0, 0, 0, 0, 0, 0, 2300, 0, 0, 0, 0, 0, 3100]
        }, {
          label: 'Факт:OK',
          type: 'bar',
          backgroundColor: '#8e5ea2',
          data: [0, 0, 0, 400, 0, 0, 0, 0, 0, 0, 0, 2305, 0, 0, 0, 0, 0]
        }
      ]
    }, {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          ticks: {
            reverse: true
          }
        }]
      }
    })
  }
}

</script>
