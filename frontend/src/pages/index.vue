<template>
    <div class="bg-bg text-text">
        <div class="hero mt-24">
            <div class="hero-content lg:flex-col flex-row grow w-full">
                <div class="h-full flex flex-col lg:flex-row justify-between w-full">
                    <div>
                        <h1 class="text-accent font-bold text-5xl">Playground</h1>
                        <h2 class="mt-4 text-secondary-content text-3xl">Feature 1</h2>
                        <p class="mt-2">Info About Feature 1</p>
                    </div>
                    <div>
                        <p class="text-lg">Title Of Card 1</p>
                        <Card />
                    </div>
                </div>
                <div class="flex flex-col lg:flex-row justify-between w-full lg:mt-20">
                    <div>
                        <h2 class="mt-4 text-secondary-content text-3xl">Feature 2</h2>
                        <p class="mt-2">Info About Feature 2</p>
                    </div>
                    <div>
                        <p class="text-lg">Title Of Card 2</p>
                        <Card />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import Card from '../components/Card.vue'
</script>