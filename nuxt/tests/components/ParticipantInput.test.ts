import { describe, it, expect } from 'vitest'
import { mountSuspended } from '@nuxt/test-utils/runtime'
import ParticipantInput from '../../components/ParticipantInput.vue'

describe('ParticipantInput', () => {
  it('should render textarea', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: [],
      },
    })

    expect(component.find('textarea').exists()).toBe(true)
  })

  it('should display participant count', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: ['Alice', 'Bob', 'Charlie'],
      },
    })

    expect(component.text()).toContain('3 名が入力されています')
  })

  it('should emit update:modelValue on input', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: [],
      },
    })

    const textarea = component.find('textarea')
    await textarea.setValue('Alice\nBob\nCharlie\nDave')

    const emitted = component.emitted('update:modelValue')
    expect(emitted).toBeTruthy()
    expect(emitted![emitted!.length - 1][0]).toEqual(['Alice', 'Bob', 'Charlie', 'Dave'])
  })

  it('should parse comma-separated input', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: [],
      },
    })

    const textarea = component.find('textarea')
    await textarea.setValue('Alice, Bob, Charlie, Dave')

    const emitted = component.emitted('update:modelValue')
    expect(emitted).toBeTruthy()
    expect(emitted![emitted!.length - 1][0]).toEqual(['Alice', 'Bob', 'Charlie', 'Dave'])
  })

  it('should show error when less than 4 participants', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: [],
      },
    })

    const textarea = component.find('textarea')
    await textarea.setValue('Alice\nBob')

    expect(component.text()).toContain('4名以上の参加者が必要です')
  })

  it('should not show error when 4 or more participants', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: [],
      },
    })

    const textarea = component.find('textarea')
    await textarea.setValue('Alice\nBob\nCharlie\nDave')

    expect(component.text()).not.toContain('4名以上の参加者が必要です')
  })

  it('should filter empty names', async () => {
    const component = await mountSuspended(ParticipantInput, {
      props: {
        modelValue: [],
      },
    })

    const textarea = component.find('textarea')
    await textarea.setValue('Alice\n\nBob\n  \nCharlie')

    const emitted = component.emitted('update:modelValue')
    expect(emitted).toBeTruthy()
    expect(emitted![emitted!.length - 1][0]).toEqual(['Alice', 'Bob', 'Charlie'])
  })
})
