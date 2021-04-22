<template>
  <v-list-item two-line>
      <v-list-item-content>
        <v-list-item-title>{{ job.InputFile.replace("\\", "/").split('/').pop() }}</v-list-item-title>
        <v-list-item-subtitle>
          {{ new Date(job.Estimation).toISOString().substr(11, 8) }}
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
    job: function(newVal, oldVal) { // watch it
      console.log('Prop changed: ', newVal.Progress, ' | was: ', oldVal.Progress)
    }
  }
  
}
</script>
