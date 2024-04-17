import React from 'react'
import { IMaskInput } from 'react-imask'

interface UUIDTextMaskProps {
  onChange: (event: { target: { name: string; value: string } }) => void
  name: string
}

export const UUIDTextMask = React.forwardRef<
  HTMLInputElement,
  UUIDTextMaskProps
>((props, ref) => {
  const { onChange, ...other } = props
  return (
    <IMaskInput
      {...other}
      mask={
        /^([*]{1})|([0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12})$/
      }
      inputRef={ref}
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      onAccept={(value: any) =>
        onChange({ target: { name: props.name, value } })
      }
      overwrite
    />
  )
})
