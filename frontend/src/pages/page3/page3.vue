<template>
<div class="title">Page 3
    <canvas id="bar-chart" width="800" height="450"></canvas>
</div>
</template>
<script>
// eslint-disable-next-line spaced-comment
//import Chart from 'Chart.js/dist/Chart.js'
export default {
  mounted () {
    console.log(document.getElementById('bar-chart')) // I'm text inside the component.
  }
}
</script>
