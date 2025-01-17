// Copyright 2024, axtlos <axtlos@disroot.org>
// SPDX-License-Identifier: GPL-3.0-ONLY

#include <stdio.h>
#include <stdlib.h>
#include <yyjson.h>
#include <string.h>

extern const char *PlugInfo()
{
  return "{\"Name\":\"shim\",\"Type\":\"1\",\"Usecontainercmds\":0}";
}

extern const char *BuildModule (char *moduleInterface, char *recipeInterface)
{
  yyjson_doc *module = yyjson_read(moduleInterface, strlen(moduleInterface), 0);
  yyjson_val *moduleRoot = yyjson_doc_get_root(module);
  yyjson_val *moduleName = yyjson_obj_get(moduleRoot, "type");
  char *name = malloc(sizeof(char)*(int)yyjson_get_len(moduleName));
  strcpy(name, yyjson_get_str(moduleName));
  yyjson_doc_free(module);

  yyjson_doc *recipe = yyjson_read(recipeInterface, strlen(recipeInterface), 0);
  yyjson_val *recipeRoot = yyjson_doc_get_root(recipe);
  yyjson_val *recipePluginPath = yyjson_obj_get(recipeRoot, "PluginPath");
  char *pluginPath = malloc(sizeof(char)*(int)yyjson_get_len(recipePluginPath));
  strcpy(pluginPath, yyjson_get_str(recipePluginPath));
  yyjson_doc_free(recipe);

  printf("[SHIM] Plugin Path: %s\n", pluginPath);
  printf("[SHIM] Name: %s\n", name);
  printf("[SHIM] Executing plugin: %s/%s\n", pluginPath, name);

  FILE *fptr;
  fptr = fopen("/tmp/meow", "w");
  fprintf(fptr, "%s\n%s", moduleInterface, recipeInterface);
  fclose(fptr);

  char *command = malloc(sizeof(char)*(strlen(pluginPath)+strlen(name))+1);
  sprintf(command, "%s/%s /tmp/meow", pluginPath, name);
  FILE *fp;
  char *path;
  fp = popen(command, "r");
  if (fp == NULL)
    return "ERROR: [SHIM] failed to launch command";

  char buffer[10];
  char *input = 0;
  size_t cur_len = 0;
  while (fgets(buffer, sizeof(buffer), fp) != NULL)
    {
      size_t buf_len = strlen(buffer);
      char *extra = realloc(input, buf_len + cur_len + 1);
      if (extra == 0)
	break;
      input = extra;
      strcpy(input + cur_len, buffer);
      cur_len += buf_len;
    }

  pclose(fp);
  return input;
}
