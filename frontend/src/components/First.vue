<template>
 <div class="container">
    <button type="button" class="btn btn-primary">{{ msg }}</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'First',
  data () {
    return {
      msg: ''
    }
  },
  methods: {
    getMessage () {
      const path = 'http://localhost:8081/test'
      axios.get(path)
        .then((res) => {
          // eslint-disable-next-line
          this.msg = 'hi '+ res.data
        })
        .catch((error) => {
          // eslint-выключение следующей строки
          console.error(error)
        })
    }
  },
  created () {
    this.getMessage()
  }

}
</script>
