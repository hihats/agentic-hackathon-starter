import { describe, it, expect } from 'vitest'
import { mountSuspended } from '@nuxt/test-utils/runtime'
import GroupSettings from '../../components/GroupSettings.vue'

describe('GroupSettings', () => {
  it('should render radio buttons', async () => {
    const component = await mountSuspended(GroupSettings)

    const radios = component.findAll('input[type="radio"]')
    expect(radios.length).toBe(2)
  })

  it('should default to group_size mode', async () => {
    const component = await mountSuspended(GroupSettings)

    const groupSizeRadio = component.find('input[value="group_size"]')
    expect((groupSizeRadio.element as HTMLInputElement).checked).toBe(true)
  })

  it('should emit groupSize on mount', async () => {
    const component = await mountSuspended(GroupSettings)

    // Wait for onMounted to execute
    await component.vm.$nextTick()

    const groupSizeEmitted = component.emitted('update:groupSize')
    expect(groupSizeEmitted).toBeTruthy()
    expect(groupSizeEmitted![0][0]).toBe(3) // default value
  })

  it('should emit numGroups when switching mode', async () => {
    const component = await mountSuspended(GroupSettings)

    const numGroupsRadio = component.find('input[value="num_groups"]')
    await numGroupsRadio.setValue(true)

    const numGroupsEmitted = component.emitted('update:numGroups')
    expect(numGroupsEmitted).toBeTruthy()
  })

  it('should show group size input in group_size mode', async () => {
    const component = await mountSuspended(GroupSettings)

    expect(component.text()).toContain('1グループあたりの人数')
  })

  it('should show num groups input in num_groups mode', async () => {
    const component = await mountSuspended(GroupSettings)

    const numGroupsRadio = component.find('input[value="num_groups"]')
    await numGroupsRadio.setValue(true)

    expect(component.text()).toContain('グループ数')
  })

  it('should emit updated groupSize value', async () => {
    const component = await mountSuspended(GroupSettings)

    const numberInput = component.find('input[type="number"]')
    await numberInput.setValue(5)

    const emitted = component.emitted('update:groupSize')
    expect(emitted).toBeTruthy()
    // Find the emission after the value change
    const lastEmission = emitted![emitted!.length - 1][0]
    expect(lastEmission).toBe(5)
  })
})
