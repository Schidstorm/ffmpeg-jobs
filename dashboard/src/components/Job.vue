<template>
  <v-list-item two-line>
      <v-list-item-content>
        <v-list-item-title>{{ job.InputFile.replace("\\", "/").split('/').pop() }}</v-list-item-title>
        <v-list-item-subtitle>
          {{ new Date(estimation).toISOString().substr(11, 8) }}
        </v-list-item-subtitle>
        <v-progress-linear :value="job.Progress * 100"></v-progress-linear>
      </v-list-item-content>
      <v-list-item-action>
          <v-btn @click="deleteJob()" icon>
            <v-icon color="red" >delete</v-icon>
          </v-btn>
        </v-list-item-action>
    </v-list-item>
</template>

<script>
export default {
  name: 'Job',
  props: {
    apiServer: String,
    job: {
      ID: Number,
      InputFile: String,
      OutputFile: String,
      Progress: Number
    }
  },
  data: () => ({
    estimation: 0
  }),
  methods: {
    deleteJob() {
      fetch(`${this.$props.apiServer}/job/${this.$props.job.ID}`, {
          method: 'DELETE'
      }).then(() => {
        // destroy the vue listeners, etc
        this.$destroy();

        // remove the element from the DOM
        this.$el.parentNode.removeChild(this.$el);
      })
    }
  },
  watch: { 
    job: function(newVal) { // watch it
      const lastChange = this.$data.lastChange || {time: Date.now(), progress: newVal.Progress}
      const change = {time: Date.now(), progress: newVal.Progress}
      const estimation = Math.floor((change.time - lastChange.time) / (change.progress - lastChange.progress))
      this.$data.lastChange = change
      this.$data.estimation = Number.isNaN(estimation) || !Number.isFinite(estimation) ? 0 : estimation
    }
  }
  
}
</script>
