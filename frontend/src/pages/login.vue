<template>
    <div class="hero h-lvh overflow-hidden mt-24">
        <div class="hero-content h-full block">
            <h1 class="w-fit">Login to your account</h1>
            <form @submit.prevent="sendLogin" class="form-control bg-secondary p-4 rounded-md w-fit">
                <div class="flex flex-col">
                    <label class="text-neutral">Username</label>
                    <input required type="input" v-model="username" class="input text-accent" placeholder="Username">
                </div>
                <div class="flex flex-col mt-4">
                    <label class="text-neutral">Password</label>
                    <input required type="password" v-model="password" class="input text-accent" placeholder="Password">
                </div>
                <button type="submit" class="btn btn-accent mt-4">Log In</button>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter();

const username = ref('')
const password = ref('')

async function sendLogin(_: Event) {
    // target: POST /api/auth
    // body: application/x-www-urlencoded
    const form = new URLSearchParams()
    form.set("username", username.value)
    form.set("password", password.value)
    const response = await fetch(import.meta.env.VITE_API_URL + '/api/auth', {
        body: form,
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })

    try {
        const output: {
            code: number;
            message: string;
        } = await response.json()
        if (response.status != 200) {
            console.log(output)
            return
        }

        // router.push('/account')
    } catch (error: unknown) {
        const err = error as Error
        console.log(err)
    }
}
</script>