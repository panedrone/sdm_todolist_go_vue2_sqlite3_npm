import Vue from 'vue';
import whoiam from "./components/whoiam.vue";
import projects from "./components/projects.vue";
import project_details from "./components/project_details.vue";
import task_details from "./components/task_details.vue";
import my_event_bus from "./components/my_event_bus";

new Vue({
    el: "#app",
    components: {
        whoiam,
        projects,
        project_details,
        task_details
    },
    methods: {
    },
    created() {
    },
    updated() {
    },
    mounted() { // https://codepen.io/g2g/pen/mdyeoXB
        my_event_bus.renderWhoIAm();
        my_event_bus.renderProjects();
    },
})