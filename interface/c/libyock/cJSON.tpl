// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#ifndef cJSON__h
#define cJSON__h

#define CJSON_STDCALL __stdcall
#define CJSON_PUBLIC(type) __declspec(dllexport) type CJSON_STDCALL

typedef struct cJSON
{
} cJSON;

/* raw json */
CJSON_PUBLIC(cJSON *)
cJSON_CreateObject(void);

/* Helper functions for creating and adding items to an object at the same time.
 * They return the added item or NULL on failure. */
CJSON_PUBLIC(cJSON *)
cJSON_AddStringToObject(cJSON *const object, const char *const name, const char *const string);

/* Render a cJSON entity to text for transfer/storage without any formatting. */
CJSON_PUBLIC(char *)
cJSON_PrintUnformatted(const cJSON *item);

/* Memory Management: the caller is always responsible to free the results from all variants of cJSON_Parse (with cJSON_Delete) and cJSON_Print (with stdlib free, cJSON_Hooks.free_fn, or cJSON_free as appropriate). The exception is cJSON_PrintPreallocated, where the caller has full responsibility of the buffer. */
/* Supply a block of JSON, and this returns a cJSON object you can interrogate. */
CJSON_PUBLIC(cJSON *)
cJSON_Parse(const char *value);

/* Check item type and return its value */
CJSON_PUBLIC(char *)
cJSON_GetStringValue(const cJSON *const item);

/* Get item "string" from object. Case insensitive. */
CJSON_PUBLIC(cJSON *)
cJSON_GetObjectItem(const cJSON *const object, const char *const string);
#endif