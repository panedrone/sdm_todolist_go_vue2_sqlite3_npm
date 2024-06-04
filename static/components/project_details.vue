<script>

import * as shared from './shared.js'
import my_event_bus from "./my_event_bus";

const NO_PROJECT = {"p_id": -1, "p_name": null, "p_tasks_count": -1}
const TASK = {"t_id": -1, "t_date": null, "t_subject": null, "t_priority": -1}

export default {
  data() {
    return {
      current_project: NO_PROJECT,
      tasks: [TASK],
      t_subject: null,
      show_project_details: false,
    }
  },
  created() {
    my_event_bus.hideProjectDetails = this.hideProjectDetails;
    my_event_bus.renderProjectDetails = this.renderProjectDetails;
    my_event_bus.renderProjectTasks = this.renderProjectTasks;
  },
  methods: {
    hideProjectDetails() {
      my_event_bus.hideTaskDetails()
      this.show_project_details = false
    },
    renderProjectDetails(p_id) {
      this.renderCurrentProject(p_id)
      this.renderProjectTasks(p_id)
      this.show_project_details = true
    },
    renderCurrentProject(p_id) {
      fetch("/api/projects/" + p_id)
          .then(async (resp) => {
            if (resp.status === 200) {
              this.current_project = await resp.json()
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    renderProjectTasks(p_id) {
      fetch("/api/projects/" + p_id + "/tasks")
          .then(async (resp) => {
            if (resp.status === 200) {
              this.tasks = await resp.json()
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    projectUpdate() {
      let p_id = this.current_project.p_id
      let json = JSON.stringify(this.current_project)
      fetch("/api/projects/" + p_id, {
        method: 'put',
        headers: shared.JSON_HEADERS,
        body: json
      })
          .then(async (resp) => {
            if (resp.status === 200) {
              my_event_bus.renderProjects();
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    projectDelete() {
      let p_id = this.current_project.p_id
      fetch("/api/projects/" + p_id, {
        method: 'delete'
      })
          .then(async (resp) => {
            if (resp.status === 204) {
              this.hideProjectDetails()
              my_event_bus.renderProjects();
            } else {
              let j = await resp.text()
              alert(resp.status + "\n" + j);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    taskCreate() {
      let p_id = this.current_project.p_id
      let json = JSON.stringify({"t_subject": this.t_subject})
      fetch("/api/projects/" + p_id + "/tasks", {
        method: 'post',
        headers: shared.JSON_HEADERS,
        body: json
      })
          .then(async (resp) => {
            if (resp.status === 201) {
              my_event_bus.renderProjects(); // update tasks count
              this.renderProjectDetails(p_id);
            } else {
              let text = await resp.text()
              alert(resp.status + "\n" + text);
            }
          })
          .catch((reason) => {
            console.log(reason)
          })
    },
    renderTaskDetails(t_id) {
      // this.$root.renderTaskDetails(t_id);
      my_event_bus.renderTaskDetails(t_id)
    }
  }
}
</script>

<template>
  <table v-if="show_project_details">
    <tr>
      <td>
        <table class="controls">
          <tr>
            <td>
              <label>
                <input type="text" class="project-header" v-model="current_project.p_name"/>
              </label>
            </td>
            <td class="w1">
              <a href="#" @click="projectUpdate()">
                <input type="button" value="✓"/>
              </a>
            </td>
            <td class="w1">
              <a href="#" @click="projectDelete()">
                <input type="button" value="x"/>
              </a>
            </td>
            <td class="w1">
              <a href="#" @click="hideProjectDetails()">
                <input type="button" value="&lt;"/>
              </a>
            </td>
          </tr>
        </table>
      </td>
    </tr>
    <tr>
      <td>
        <table class="section">
          <tr>
            <th>Date</th>
            <th>Subject</th>
            <th>Priority</th>
          </tr>
          <tr v-for="t in tasks">
            <td>
              {{ t.t_date }}
            </td>
            <td>
              <a href="#" @click="renderTaskDetails(t.t_id)">
                {{ t.t_subject }}
              </a>
            </td>
            <td style="text-align: center;">
              {{ t.t_priority }}
            </td>
          </tr>
        </table>
      </td>
    </tr>
    <tr>
      <td>
        <table class="controls">
          <tr>
            <td>
              <label>
                <input type="text" v-model="t_subject"/>
              </label>
            </td>
            <td class="w1">
              <a href="#" @click="taskCreate()">
                <input type="button" value="+"/>
              </a>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</template>

<style scoped>

</style>