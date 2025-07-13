<script lang="ts" setup>
    import { reactive } from "vue";
    import { Icon } from "@iconify/vue";
    import { useColorMode } from "@vueuse/core";
    import { Button } from "@/components/ui/button";
    import {
        DropdownMenu,
        DropdownMenuContent,
        DropdownMenuItem,
        DropdownMenuTrigger,
    } from "@/components/ui/dropdown-menu";
    import { Greet } from "../../wailsjs/go/main/App";

    const mode = useColorMode();

    const data = reactive({
        name: "",
        resultText: "Please enter your name below 👇",
    });

    function greet() {
        Greet(data.name).then((result) => {
            data.resultText = result;
        });
    }
</script>

<template>
    <main>
        <h1 class="bg-blue-800 font-bold text-gray-200">Tailwind Test</h1>
        <div>
            <Button variant="destructive">Click me</Button>
        </div>
        <div>
            <DropdownMenu>
                <DropdownMenuTrigger as-child>
                    <Button variant="outline">
                        <Icon
                            icon="radix-icons:moon"
                            class="h-[1.2rem] w-[1.2rem] scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90"
                        />
                        <Icon
                            icon="radix-icons:sun"
                            class="absolute h-[1.2rem] w-[1.2rem] scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0"
                        />
                        <span class="sr-only">Toggle theme</span>
                    </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                    <DropdownMenuItem @click="mode = 'light'">
                        Light
                    </DropdownMenuItem>
                    <DropdownMenuItem @click="mode = 'dark'">
                        Dark
                    </DropdownMenuItem>
                    <DropdownMenuItem @click="mode = 'auto'">
                        System
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
        <div id="result" class="result">{{ data.resultText }}</div>
        <div id="input" class="input-box">
            <input
                id="name"
                v-model="data.name"
                autocomplete="off"
                class="input"
                type="text"
            />
            <button class="btn" @click="greet">Greet</button>
        </div>
    </main>
</template>

<style scoped>
    .result {
        height: 20px;
        line-height: 20px;
        margin: 1.5rem auto;
    }

    .input-box .btn {
        width: 60px;
        height: 30px;
        line-height: 30px;
        border-radius: 3px;
        border: none;
        margin: 0 0 0 20px;
        padding: 0 8px;
        cursor: pointer;
    }

    .input-box .btn:hover {
        background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
        color: #333333;
    }

    .input-box .input {
        border: none;
        border-radius: 3px;
        outline: none;
        height: 30px;
        line-height: 30px;
        padding: 0 10px;
        background-color: rgba(240, 240, 240, 1);
        -webkit-font-smoothing: antialiased;
    }

    .input-box .input:hover {
        border: none;
        background-color: rgba(255, 255, 255, 1);
    }

    .input-box .input:focus {
        border: none;
        background-color: rgba(255, 255, 255, 1);
    }
</style>
