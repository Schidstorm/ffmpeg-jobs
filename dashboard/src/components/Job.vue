<template>
  <v-list-item two-line>
      <v-list-item-content>
        <v-list-item-title>{{ job.InputFile }}</v-list-item-title>
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
  }
  
}
</script>
