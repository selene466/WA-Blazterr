<script lang="ts" setup>
    import { onMounted, ref } from "vue";
    import { Button } from "@/components/ui/button";
    import { ShowWhatsapp, StartWhatsapp } from "../wailsjs/go/main/App";
    import * as runtime from "../wailsjs/runtime/runtime";

    // mounted
    onMounted(async () => {
        runtime.EventsOn("whatsappEvent", (msg: string) => {
            if (msg === "running") isWhatsappRunning.value = true;
        });
    });

    // var
    let data = "test";

    // ref
    const isWhatsappRunning = ref(false);

    // function
    const btnStartWhatsapp = async () => {
        await StartWhatsapp();
    };

    async function showWhatsapp() {
        data = await ShowWhatsapp();
    }
</script>

<template>
    <div>
        <Button @click="btnStartWhatsapp()" :disabled="isWhatsappRunning">
            Start
        </Button>
        <Button @click="showWhatsapp">Show</Button>
        <div>
            {{ data }}
        </div>
    </div>
</template>
