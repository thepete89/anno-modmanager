<script lang="ts">
  import { page } from "$app/state";
	import "../app.css";
	import favicon from "$lib/assets/favicon.svg";
	import { Navigation } from "@skeletonlabs/skeleton-svelte";
	import { HouseIcon, SettingsIcon, CastleIcon, StarIcon, LayoutList } from "@lucide/svelte";
	
	let { children } = $props();

	const links = [
		{ label: 'Mods', href: '/', icon: HouseIcon },
    { label: 'Subscriptions', href: '/', icon: StarIcon },
    { label: 'Collections', href: '/', icon: LayoutList }
	];
  let anchorRail = 'btn hover:preset-tonal aspect-square w-full max-w-[84px] flex flex-col items-center gap-0.5';
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="w-full h-svh grid grid-cols-[auto_1fr] border border-surface-200-800">
  <!-- --- -->
  <Navigation layout="rail">
    <Navigation.Header>
      <div class="aspect-square w-full max-w-[84px] flex flex-col items-center gap-0.5">
        <CastleIcon class="size-5" />
        <span class="text-xs text-center">Anno Mod Manager</span>
      </div>
    </Navigation.Header>
    <Navigation.Content>
      <Navigation.Menu>
        {#each links as link (link)}
          {@const Icon = link.icon}
          <a href={link.href} class={[anchorRail, link.href === page.url.pathname && "preset-tonal-primary"]}>
            <Icon class="size-5" />
            <span class="text-xs">{link.label}</span>
          </a>
        {/each}
      </Navigation.Menu>
    </Navigation.Content>
    <Navigation.Footer>
      <a href="/settings" class={[anchorRail, "/settings" === page.url.pathname && "preset-tonal-primary"]} title="Settings" aria-label="Settings">
        <SettingsIcon class="size-5" />
        <span class="text-xs">Settings</span>
      </a>
    </Navigation.Footer>
  </Navigation>
  <!-- --- -->
  <main class="space-y-4 p-4">
    {@render children()}
  </main>
  <!--
    <div class="flex justify-center items-center">
    </div>
  -->
</div>

