import { describe, it, expect } from 'vitest'
import { mountSuspended } from '@nuxt/test-utils/runtime'
import ResultDisplay from '../../components/ResultDisplay.vue'

describe('ResultDisplay', () => {
  it('should not render when groups is empty', async () => {
    const component = await mountSuspended(ResultDisplay, {
      props: {
        groups: [],
      },
    })

    expect(component.find('.result-display').exists()).toBe(false)
  })

  it('should render groups when provided', async () => {
    const component = await mountSuspended(ResultDisplay, {
      props: {
        groups: [
          ['Alice', 'Bob'],
          ['Charlie', 'Dave'],
        ],
      },
    })

    expect(component.find('.result-display').exists()).toBe(true)
    expect(component.text()).toContain('シャッフル結果')
  })

  it('should display correct group labels', async () => {
    const component = await mountSuspended(ResultDisplay, {
      props: {
        groups: [
          ['Alice', 'Bob'],
          ['Charlie', 'Dave'],
          ['Eve', 'Frank'],
        ],
      },
    })

    expect(component.text()).toContain('グループ A')
    expect(component.text()).toContain('グループ B')
    expect(component.text()).toContain('グループ C')
  })

  it('should display all members', async () => {
    const component = await mountSuspended(ResultDisplay, {
      props: {
        groups: [
          ['Alice', 'Bob'],
          ['Charlie', 'Dave'],
        ],
      },
    })

    expect(component.text()).toContain('Alice')
    expect(component.text()).toContain('Bob')
    expect(component.text()).toContain('Charlie')
    expect(component.text()).toContain('Dave')
  })

  it('should display member count per group', async () => {
    const component = await mountSuspended(ResultDisplay, {
      props: {
        groups: [
          ['Alice', 'Bob', 'Charlie'],
          ['Dave', 'Eve'],
        ],
      },
    })

    expect(component.text()).toContain('3名')
    expect(component.text()).toContain('2名')
  })

  it('should render correct number of group cards', async () => {
    const component = await mountSuspended(ResultDisplay, {
      props: {
        groups: [
          ['Alice'],
          ['Bob'],
          ['Charlie'],
          ['Dave'],
        ],
      },
    })

    const groupCards = component.findAll('.group')
    expect(groupCards.length).toBe(4)
  })
})
