#pragma once

#include <stddef.h>
#include <stdint.h>

typedef struct {
  uint8_t *bytes;
  size_t len;
} Bytes;