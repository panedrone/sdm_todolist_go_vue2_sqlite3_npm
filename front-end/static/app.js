import Vue from 'vue';
import whoiam from "./components/whoiam.vue";
import server_error from "./components/server_error.vue";
import projects from "./components/project_list.vue";
import project_details from "./components/project_details.vue";
import task_details from "./components/task_details.vue";
import fire from "./components/event_bus";

new Vue({
    el: "#app",
    components: {
        whoiam,
        server_error,
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
        fire.renderWhoIAm();
        fire.renderProjects();
    },
})