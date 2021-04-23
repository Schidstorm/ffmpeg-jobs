<template>
  <v-card
    class="mx-auto"
    max-width="600"
    tile
  >
    <v-card-title>
      Jobs 
    </v-card-title>

    <Job v-for="job in jobs" :key="job.ID" :apiServer="apiServer" :job="job">

    </Job>
  </v-card>
</template>

<script>
import Job from './Job.vue'

export default {
  name: 'JobList',
  props: {
    apiServer: String
  },
  data () {
    return {
      loading: false,
      post: null,
      error: null,
      jobs: []
    }
  },
  created () {
    // fetch the data when the view is created and the data is
    // already being observed
    this.fetchData()
  },
  methods: {
    fetchData () {
      this.error = this.post = null
      this.loading = true
      const url = `${this.$props.apiServer}/job`
      // replace `getPost` with your data fetching util / API wrapper
      this.loop(
        () => fetch(url)
          .catch(err => this.error = err)
          .then(response => response.json())
          .then(data => this.$data.jobs = data.Data),
      20000)
    },
    loop(callback, timeout) {
      callback()
      window.setInterval(callback, timeout)
    }
  },
  components: {
    Job
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
