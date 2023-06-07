// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#ifndef __TEST_MYLIB_H__
#define __TEST_MYLIB_H__
#include <stdio.h>

typedef char *string;
typedef struct person
{
    string name;
    int age;
    string telephone;
} Person;

#define TEST_API __declspec(dllexport)

TEST_API void hello();
TEST_API string hello2(string, int);
TEST_API Person hello3(Person);

#endif