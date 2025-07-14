<script lang="ts" setup>
    import { Button } from "@/components/ui/button";
    import { StartWhatsapp, ShowWhatsapp } from "../wailsjs/go/main/App";

    // var
    let data = "test";

    // function
    function startWhatsapp() {
        StartWhatsapp();
    }

    async function showWhatsapp() {
        data = await ShowWhatsapp();
    }
</script>

<template>
    <div>
        <Button @click="startWhatsapp">Start</Button>
        <Button @click="showWhatsapp">Show</Button>
        <div>
            {{ data }}
        </div>
    </div>
</template>
