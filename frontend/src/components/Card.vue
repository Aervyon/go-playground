<template>
    <div class="p-4 card bg-secondary max-w-[512px] my-auto text-text-2 shadow-lg">
        <div class="card-body">
            <div class="w-full items-center" :class="{
                'flex space-x-4': item.length + (user?.length ?? 14) <= 40,
                'space-y-2': item.length + (user?.length ?? 14) > 40
            }">
                <h1 class="text-accent text-xl font-bold w-fit card-title">{{ user ?? 'Anonymous User' }}</h1>
                <h2 class="font-bold badge badge-primary badge-lg text-text-2">{{ item }}</h2>
            </div>
            <div class="flex items-center">
                <p>{{ data }}</p>
                <div class="rating disabled">
                    <input disabled type="radio" :name="'rating-' + index" class="cursor-default mask mask-star bg-yellow-400" checked>
                    <input disabled type="radio" :name="'rating-' + index" class="cursor-default mask mask-star bg-yellow-400">
                    <input disabled type="radio" :name="'rating-' + index" class="cursor-default mask mask-star bg-yellow-400">
                    <input disabled type="radio" :name="'rating-' + index" class="cursor-default mask mask-star bg-yellow-400">
                    <input disabled type="radio" :name="'rating-' + index" class="cursor-default mask mask-star bg-yellow-400">
                </div>
            </div>
            <div class="flex items-center">
                <p>{{ data2 }}</p>
                <div class="rating">
                    <input disabled type="radio" :name="'rating-' + (index + 1)" class="cursor-default mask mask-star bg-yellow-400">
                    <input disabled type="radio" :name="'rating-' + (index + 1)" class="cursor-default mask mask-star bg-yellow-400">
                    <input disabled type="radio" :name="'rating-' + (index + 1)" class="cursor-default mask mask-star bg-yellow-400" checked>
                    <input disabled type="radio" :name="'rating-' + (index + 1)" class="cursor-default mask mask-star bg-yellow-400">
                    <input disabled type="radio" :name="'rating-' + (index + 1)" class="cursor-default mask mask-star bg-yellow-400">
                </div>
            </div>
            <p class="text-text">{{ info }}</p>
            <p>Footnotes</p>
        </div>
    </div>
</template>

<script setup lang="ts">
defineProps<{
    index: number;
    item: string;
    user?: string;
    data: string;
    data2: string;
    info: string;
}>()
</script>