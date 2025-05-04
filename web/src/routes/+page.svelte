<script lang="ts">
    import { user } from '$lib/stores/auth';
    import { goto } from '$app/navigation';
</script>

<div class="text-center py-12">
    <h1 class="text-4xl font-bold mb-6">Welcome to Events App</h1>
    {#if $user}
        <p class="text-xl mb-4">Hello, {$user.email}!</p>
        <div class="space-x-4">
            <button 
                on:click={() => goto('/events')}
                class="bg-blue-500 hover:bg-blue-600 text-white px-6 py-2 rounded transition-colors"
                style="background-color: #3b82f6; color: white; padding: 0.5rem 1.5rem; border-radius: 0.25rem;"
            >
                View All Events
            </button>
            <button 
                on:click={() => goto('/my-events')}
                class="bg-green-500 hover:bg-green-600 text-white px-6 py-2 rounded transition-colors"
                style="background-color: #22c55e; color: white; padding: 0.5rem 1.5rem; border-radius: 0.25rem;"
            >
                My Events
            </button>
        </div>
    {:else}
        <p class="text-xl mb-4">Please login or sign up to view and register for events</p>
        <div class="space-x-4">
            <button 
                on:click={() => goto('/login')}
                class="bg-blue-500 hover:bg-blue-600 text-white px-6 py-2 rounded transition-colors"
                style="background-color: #3b82f6; color: white; padding: 0.5rem 1.5rem; border-radius: 0.25rem;"
            >
                Login
            </button>
            <button 
                on:click={() => goto('/signup')}
                class="bg-green-500 hover:bg-green-600 text-white px-6 py-2 rounded transition-colors"
                style="background-color: #22c55e; color: white; padding: 0.5rem 1.5rem; border-radius: 0.25rem;"
            >
                Sign Up
            </button>
        </div>
    {/if}
</div>
