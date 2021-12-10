<template>
  <div>
    <p>hello shoezoo!!</p>
    <button @click="ping">ping</button>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { createRequestClient } from '~/store/request-client'

export default Vue.extend({
  methods: {
    ping() {
      const req = createRequestClient(this.$axios)
      req.get('http://localhost:8080/ping').then((res) => {
        console.log(res)
      })
    },
  },
})
</script>
