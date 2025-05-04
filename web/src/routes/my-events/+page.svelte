<script lang="ts">
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
    import { user } from '$lib/stores/auth';
    import { getRegisteredEvents, cancelRegistration } from '$lib/services/api';

    let events: any[] = [];
    let loading = true;
    let error = '';

    onMount(async () => {
        if (!$user) {
            goto('/login');
            return;
        }

        try {
            events = await getRegisteredEvents();
        } catch (e) {
            error = 'Failed to load registered events';
        } finally {
            loading = false;
        }
    });

    async function handleCancel(eventId: number) {
        try {
            await cancelRegistration(eventId);
            // Refresh events list
            events = await getRegisteredEvents();
        } catch (e) {
            error = 'Failed to cancel registration';
        }
    }
</script>

<div class="max-w-4xl mx-auto">
    <h1 class="text-3xl font-bold mb-6">My Events</h1>

    {#if error}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {error}
        </div>
    {/if}

    {#if loading}
        <div class="text-center py-8">Loading events...</div>
    {:else if events.length === 0}
        <div class="text-center py-8">
            <p class="text-gray-600">You haven't registered for any events yet.</p>
            <a href="/events" class="text-blue-500 hover:text-blue-600 mt-2 inline-block">
                Browse all events
            </a>
        </div>
    {:else}
        <div class="grid gap-6">
            {#each events as event}
                <div class="bg-white rounded-lg shadow-md p-6">
                    <h2 class="text-xl font-bold mb-2">{event.name}</h2>
                    <p class="text-gray-600 mb-2">{event.description}</p>
                    <p class="text-gray-600 mb-2">
                        <strong>Date:</strong> {new Date(event.date_time).toLocaleString()}
                    </p>
                    <p class="text-gray-600 mb-4">
                        <strong>Location:</strong> {event.location}
                    </p>
                    <button
                        on:click={() => handleCancel(event.id)}
                        class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
                    >
                        Cancel Registration
                    </button>
                </div>
            {/each}
        </div>
    {/if}
</div> 