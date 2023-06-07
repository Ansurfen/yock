// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#include "mylib.h"

TEST_API void hello()
{
    printf("Hello World!\n");
}

TEST_API string hello2(string name, int age)
{
    printf("Hello %s, and your age is %d?\n", name, age);
    return "Hello Yock";
}

TEST_API Person hello3(Person p)
{
    printf("name: %s, age: %d, telephone: %s\n", p.name, p.age, p.telephone);
    p.name = "yock";
    p.age = -1;
    p.telephone = "000000";
    return p;
}