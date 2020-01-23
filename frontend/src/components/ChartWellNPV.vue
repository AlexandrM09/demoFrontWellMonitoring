<script>
// import { Line } from 'vue-chartjs'
import {Doughnut} from 'vue-chartjs'
// const { reactiveProp } = mixins

export default {
  extends: Doughnut,
  // mixins: [reactiveProp],
  props: ['options'],
  mounted () {
    this.renderChart({
      labels: ['Проходка', 'ПВР', 'Крепление', 'Ремонт', 'Простои', 'Аварии'],
      datasets: [
        {
          label: 'Баланс времени',
          backgroundColor: ['#3e95cd', '#8e5ea2', '#3cba9f', '#e8c3b9', '#c45850', 'red'],
          data: [43, 30, 15, 7, 5, 7]
        }
      ]
    }, {
      title: {
        display: true,
        text: 'Баланс времени'
      }
    })
  }
}

</script>
