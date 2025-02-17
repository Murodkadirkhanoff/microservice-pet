<script setup>
import AppLayout from "@/Layouts/AppLayout.vue";
import Welcome from "@/Components/Welcome.vue";
</script>

<template>
    <AppLayout title="Test">
        <template #header>
            <h2 class="font-semibold text-xl text-gray-800 leading-tight">
                Projects
            </h2>
        </template>

        <main
            class="flex flex-1 flex-col pb-2 lg:min-w-0 lg:pt-2 lg:pr-2 lg:pl-64"
        >
            <div
                class="grow p-6 lg:rounded-lg lg:bg-white lg:p-10 lg:ring-1 lg:shadow-xs lg:ring-zinc-950/5 dark:lg:bg-zinc-900 dark:lg:ring-white/10"
            >
                <div class="mx-auto max-w-6xl">
                    <div class="flex flex-wrap items-end justify-between gap-4">
                        <div class="max-sm:w-full sm:flex-1">
                            <h1
                                class="text-2xl/8 font-semibold text-zinc-950 sm:text-xl/8 dark:text-white"
                            >
                                Projects
                            </h1>
                        </div>
                        <button
                            class="relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:outline-hidden data-focus:outline data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-blue-500 data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent bg-(--btn-border) dark:bg-(--btn-bg) before:absolute before:inset-0 before:-z-10 before:rounded-[calc(var(--radius-lg)-1px)] before:bg-(--btn-bg) before:shadow-sm dark:before:hidden dark:border-white/5 after:absolute after:inset-0 after:-z-10 after:rounded-[calc(var(--radius-lg)-1px)] after:shadow-[shadow:inset_0_1px_--theme(--color-white/15%)] data-active:after:bg-(--btn-hover-overlay) data-hover:after:bg-(--btn-hover-overlay) dark:after:-inset-px dark:after:rounded-lg data-disabled:before:shadow-none data-disabled:after:shadow-none text-white [--btn-bg:var(--color-zinc-900)] [--btn-border:var(--color-zinc-950)]/90 [--btn-hover-overlay:var(--color-white)]/10 dark:text-white dark:[--btn-bg:var(--color-zinc-600)] dark:[--btn-hover-overlay:var(--color-white)]/5 [--btn-icon:var(--color-zinc-400)] data-active:[--btn-icon:var(--color-zinc-300)] data-hover:[--btn-icon:var(--color-zinc-300)] cursor-default"
                            type="button"
                            data-headlessui-state=""
                        >
                            <span
                                class="absolute top-1/2 left-1/2 size-[max(100%,2.75rem)] -translate-x-1/2 -translate-y-1/2 [@media(pointer:fine)]:hidden"
                                aria-hidden="true"
                            ></span
                            >Create event
                        </button>
                    </div>
                    <ul class="mt-10">
                        <li>
                            <hr
                                role="presentation"
                                class="w-full border-t border-zinc-950/10 dark:border-white/10"
                            />
                            <div class="flex items-center justify-between">
                                <div class="flex gap-6 py-6">
                                    <div class="w-32 shrink-0">
                                        <a
                                            aria-hidden="true"
                                            data-headlessui-state=""
                                            href="/events/1000"
                                            ><img
                                                class="aspect-3/2 rounded-lg shadow-sm"
                                                src="/events/bear-hug.jpg"
                                                alt=""
                                        /></a>
                                    </div>
                                    <div class="space-y-1.5">
                                        <div class="text-base/6 font-semibold">
                                            <a
                                                data-headlessui-state=""
                                                href="/events/1000"
                                                >Bear Hug: Live in Concert</a
                                            >
                                        </div>
                                        <div class="text-xs/6 text-zinc-500">
                                            May 20, 2024 at 10 PM
                                            <span aria-hidden="true">·</span>
                                            Harmony Theater, Winnipeg, MB
                                        </div>
                                        <div class="text-xs/6 text-zinc-600">
                                            350/500 tickets sold
                                        </div>
                                    </div>
                                </div>
                                <div class="flex items-center gap-4">
                                    <span
                                        class="max-sm:hidden inline-flex items-center gap-x-1.5 rounded-md px-1.5 py-0.5 text-sm/5 font-medium sm:text-xs/5 forced-colors:outline bg-lime-400/20 text-lime-700 group-data-hover:bg-lime-400/30 dark:bg-lime-400/10 dark:text-lime-300 dark:group-data-hover:bg-lime-400/15"
                                        >On Sale</span
                                    ><button
                                        aria-label="More options"
                                        id="headlessui-menu-button-:r2p:"
                                        type="button"
                                        aria-haspopup="menu"
                                        aria-expanded="false"
                                        data-headlessui-state=""
                                        class="relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:outline-hidden data-focus:outline data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-blue-500 data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent text-zinc-950 data-active:bg-zinc-950/5 data-hover:bg-zinc-950/5 dark:text-white dark:data-active:bg-white/10 dark:data-hover:bg-white/10 [--btn-icon:var(--color-zinc-500)] data-active:[--btn-icon:var(--color-zinc-700)] data-hover:[--btn-icon:var(--color-zinc-700)] dark:[--btn-icon:var(--color-zinc-500)] dark:data-active:[--btn-icon:var(--color-zinc-400)] dark:data-hover:[--btn-icon:var(--color-zinc-400)] cursor-default"
                                    >
                                        <span
                                            class="absolute top-1/2 left-1/2 size-[max(100%,2.75rem)] -translate-x-1/2 -translate-y-1/2 [@media(pointer:fine)]:hidden"
                                            aria-hidden="true"
                                        ></span
                                        ><svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 16 16"
                                            fill="currentColor"
                                            aria-hidden="true"
                                            data-slot="icon"
                                        >
                                            <path
                                                d="M8 2a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM8 6.5a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM9.5 12.5a1.5 1.5 0 1 0-3 0 1.5 1.5 0 0 0 3 0Z"
                                            ></path>
                                        </svg>
                                    </button>
                                </div>
                            </div>
                        </li>
                        <li>
                            <hr
                                role="presentation"
                                class="w-full border-t border-zinc-950/5 dark:border-white/5"
                            />
                            <div class="flex items-center justify-between">
                                <div class="flex gap-6 py-6">
                                    <div class="w-32 shrink-0">
                                        <a
                                            aria-hidden="true"
                                            data-headlessui-state=""
                                            href="/events/1001"
                                            ><img
                                                class="aspect-3/2 rounded-lg shadow-sm"
                                                src="/events/six-fingers.jpg"
                                                alt=""
                                        /></a>
                                    </div>
                                    <div class="space-y-1.5">
                                        <div class="text-base/6 font-semibold">
                                            <a
                                                data-headlessui-state=""
                                                href="/events/1001"
                                                >Six Fingers — DJ Set</a
                                            >
                                        </div>
                                        <div class="text-xs/6 text-zinc-500">
                                            Jun 2, 2024 at 8 PM
                                            <span aria-hidden="true">·</span>
                                            Moonbeam Arena, Uxbridge, ON
                                        </div>
                                        <div class="text-xs/6 text-zinc-600">
                                            72/150 tickets sold
                                        </div>
                                    </div>
                                </div>
                                <div class="flex items-center gap-4">
                                    <span
                                        class="max-sm:hidden inline-flex items-center gap-x-1.5 rounded-md px-1.5 py-0.5 text-sm/5 font-medium sm:text-xs/5 forced-colors:outline bg-lime-400/20 text-lime-700 group-data-hover:bg-lime-400/30 dark:bg-lime-400/10 dark:text-lime-300 dark:group-data-hover:bg-lime-400/15"
                                        >On Sale</span
                                    ><button
                                        aria-label="More options"
                                        id="headlessui-menu-button-:r2v:"
                                        type="button"
                                        aria-haspopup="menu"
                                        aria-expanded="false"
                                        data-headlessui-state=""
                                        class="relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:outline-hidden data-focus:outline data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-blue-500 data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent text-zinc-950 data-active:bg-zinc-950/5 data-hover:bg-zinc-950/5 dark:text-white dark:data-active:bg-white/10 dark:data-hover:bg-white/10 [--btn-icon:var(--color-zinc-500)] data-active:[--btn-icon:var(--color-zinc-700)] data-hover:[--btn-icon:var(--color-zinc-700)] dark:[--btn-icon:var(--color-zinc-500)] dark:data-active:[--btn-icon:var(--color-zinc-400)] dark:data-hover:[--btn-icon:var(--color-zinc-400)] cursor-default"
                                    >
                                        <span
                                            class="absolute top-1/2 left-1/2 size-[max(100%,2.75rem)] -translate-x-1/2 -translate-y-1/2 [@media(pointer:fine)]:hidden"
                                            aria-hidden="true"
                                        ></span
                                        ><svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 16 16"
                                            fill="currentColor"
                                            aria-hidden="true"
                                            data-slot="icon"
                                        >
                                            <path
                                                d="M8 2a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM8 6.5a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM9.5 12.5a1.5 1.5 0 1 0-3 0 1.5 1.5 0 0 0 3 0Z"
                                            ></path>
                                        </svg>
                                    </button>
                                </div>
                            </div>
                        </li>
                        <li>
                            <hr
                                role="presentation"
                                class="w-full border-t border-zinc-950/5 dark:border-white/5"
                            />
                            <div class="flex items-center justify-between">
                                <div class="flex gap-6 py-6">
                                    <div class="w-32 shrink-0">
                                        <a
                                            aria-hidden="true"
                                            data-headlessui-state=""
                                            href="/events/1002"
                                            ><img
                                                class="aspect-3/2 rounded-lg shadow-sm"
                                                src="/events/we-all-look-the-same.jpg"
                                                alt=""
                                        /></a>
                                    </div>
                                    <div class="space-y-1.5">
                                        <div class="text-base/6 font-semibold">
                                            <a
                                                data-headlessui-state=""
                                                href="/events/1002"
                                                >We All Look The Same</a
                                            >
                                        </div>
                                        <div class="text-xs/6 text-zinc-500">
                                            Aug 5, 2024 at 4 PM
                                            <span aria-hidden="true">·</span>
                                            Electric Coliseum, New York, NY
                                        </div>
                                        <div class="text-xs/6 text-zinc-600">
                                            275/275 tickets sold
                                        </div>
                                    </div>
                                </div>
                                <div class="flex items-center gap-4">
                                    <span
                                        class="max-sm:hidden inline-flex items-center gap-x-1.5 rounded-md px-1.5 py-0.5 text-sm/5 font-medium sm:text-xs/5 forced-colors:outline bg-zinc-600/10 text-zinc-700 group-data-hover:bg-zinc-600/20 dark:bg-white/5 dark:text-zinc-400 dark:group-data-hover:bg-white/10"
                                        >Closed</span
                                    ><button
                                        aria-label="More options"
                                        id="headlessui-menu-button-:r35:"
                                        type="button"
                                        aria-haspopup="menu"
                                        aria-expanded="false"
                                        data-headlessui-state=""
                                        class="relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:outline-hidden data-focus:outline data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-blue-500 data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent text-zinc-950 data-active:bg-zinc-950/5 data-hover:bg-zinc-950/5 dark:text-white dark:data-active:bg-white/10 dark:data-hover:bg-white/10 [--btn-icon:var(--color-zinc-500)] data-active:[--btn-icon:var(--color-zinc-700)] data-hover:[--btn-icon:var(--color-zinc-700)] dark:[--btn-icon:var(--color-zinc-500)] dark:data-active:[--btn-icon:var(--color-zinc-400)] dark:data-hover:[--btn-icon:var(--color-zinc-400)] cursor-default"
                                    >
                                        <span
                                            class="absolute top-1/2 left-1/2 size-[max(100%,2.75rem)] -translate-x-1/2 -translate-y-1/2 [@media(pointer:fine)]:hidden"
                                            aria-hidden="true"
                                        ></span
                                        ><svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 16 16"
                                            fill="currentColor"
                                            aria-hidden="true"
                                            data-slot="icon"
                                        >
                                            <path
                                                d="M8 2a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM8 6.5a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM9.5 12.5a1.5 1.5 0 1 0-3 0 1.5 1.5 0 0 0 3 0Z"
                                            ></path>
                                        </svg>
                                    </button>
                                </div>
                            </div>
                        </li>
                        <li>
                            <hr
                                role="presentation"
                                class="w-full border-t border-zinc-950/5 dark:border-white/5"
                            />
                            <div class="flex items-center justify-between">
                                <div class="flex gap-6 py-6">
                                    <div class="w-32 shrink-0">
                                        <a
                                            aria-hidden="true"
                                            data-headlessui-state=""
                                            href="/events/1003"
                                            ><img
                                                class="aspect-3/2 rounded-lg shadow-sm"
                                                src="/events/viking-people.jpg"
                                                alt=""
                                        /></a>
                                    </div>
                                    <div class="space-y-1.5">
                                        <div class="text-base/6 font-semibold">
                                            <a
                                                data-headlessui-state=""
                                                href="/events/1003"
                                                >Viking People</a
                                            >
                                        </div>
                                        <div class="text-xs/6 text-zinc-500">
                                            Dec 31, 2024 at 8 PM
                                            <span aria-hidden="true">·</span>
                                            Tapestry Hall, Cambridge, ON
                                        </div>
                                        <div class="text-xs/6 text-zinc-600">
                                            6/40 tickets sold
                                        </div>
                                    </div>
                                </div>
                                <div class="flex items-center gap-4">
                                    <span
                                        class="max-sm:hidden inline-flex items-center gap-x-1.5 rounded-md px-1.5 py-0.5 text-sm/5 font-medium sm:text-xs/5 forced-colors:outline bg-lime-400/20 text-lime-700 group-data-hover:bg-lime-400/30 dark:bg-lime-400/10 dark:text-lime-300 dark:group-data-hover:bg-lime-400/15"
                                        >On Sale</span
                                    ><button
                                        aria-label="More options"
                                        id="headlessui-menu-button-:r3b:"
                                        type="button"
                                        aria-haspopup="menu"
                                        aria-expanded="false"
                                        data-headlessui-state=""
                                        class="relative isolate inline-flex items-baseline justify-center gap-x-2 rounded-lg border text-base/6 font-semibold px-[calc(--spacing(3.5)-1px)] py-[calc(--spacing(2.5)-1px)] sm:px-[calc(--spacing(3)-1px)] sm:py-[calc(--spacing(1.5)-1px)] sm:text-sm/6 focus:outline-hidden data-focus:outline data-focus:outline-2 data-focus:outline-offset-2 data-focus:outline-blue-500 data-disabled:opacity-50 *:data-[slot=icon]:-mx-0.5 *:data-[slot=icon]:my-0.5 *:data-[slot=icon]:size-5 *:data-[slot=icon]:shrink-0 *:data-[slot=icon]:self-center *:data-[slot=icon]:text-(--btn-icon) sm:*:data-[slot=icon]:my-1 sm:*:data-[slot=icon]:size-4 forced-colors:[--btn-icon:ButtonText] forced-colors:data-hover:[--btn-icon:ButtonText] border-transparent text-zinc-950 data-active:bg-zinc-950/5 data-hover:bg-zinc-950/5 dark:text-white dark:data-active:bg-white/10 dark:data-hover:bg-white/10 [--btn-icon:var(--color-zinc-500)] data-active:[--btn-icon:var(--color-zinc-700)] data-hover:[--btn-icon:var(--color-zinc-700)] dark:[--btn-icon:var(--color-zinc-500)] dark:data-active:[--btn-icon:var(--color-zinc-400)] dark:data-hover:[--btn-icon:var(--color-zinc-400)] cursor-default"
                                    >
                                        <span
                                            class="absolute top-1/2 left-1/2 size-[max(100%,2.75rem)] -translate-x-1/2 -translate-y-1/2 [@media(pointer:fine)]:hidden"
                                            aria-hidden="true"
                                        ></span
                                        ><svg
                                            xmlns="http://www.w3.org/2000/svg"
                                            viewBox="0 0 16 16"
                                            fill="currentColor"
                                            aria-hidden="true"
                                            data-slot="icon"
                                        >
                                            <path
                                                d="M8 2a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM8 6.5a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM9.5 12.5a1.5 1.5 0 1 0-3 0 1.5 1.5 0 0 0 3 0Z"
                                            ></path>
                                        </svg>
                                    </button>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </main>
    </AppLayout>
</template>
