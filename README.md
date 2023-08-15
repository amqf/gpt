# About

This is a simple client for GPT Chat used in terminal for general purpouse.

# Installation on Ubuntu

```bash
$ ./ci/build.sh
$ ./ci/package.sh
$ sudo dpkg -i ./release/gpt-beta.deb
```

# Generic installation
```bash
$ ./ci/build.sh
$ sudo mv ./usr/bin /usr/bin/gpt-beta
$ sudo chmod +x /usr/bin/gpt-beta
```


# Usage

First set your API_KEY.
```bash
$ export OPENAI_API_KEY=<your_openai_api_key>
```

Then do your prompt and use default temperature 0.7:
```bash
$ gpt-beta -prompt="What can GPT Chat 3.5 do?"
```

Or do your prompt with custom temperature with 0, 1 or a number between both
```bash
$ gpt-beta -prompt="What can GPT Chat 3.5 do?" -temperature=0.7
```