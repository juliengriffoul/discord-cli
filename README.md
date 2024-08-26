# discord-cli
Golang CLI to send status Discord messages.

## Build
```
go build -o bin/
```
## Usage
```
./bin/discord-cli -t <discord-token> -c <discord-channel-id> -m "Message to send" -f <bool>
```
Flag *-f* allow to specify if message should notify a failure (defaults to *false*: success notification).  
In a Github Actions workflow:
```yml
    - name: Notify success
      if: ${{ success() }}
      run: |
        docker run --rm ghcr.io/juliengriffoul/discord-cli:latest \
          -t "${{ secrets.DISCORD_TOKEN }}" \
          -c ${{ env.DISCORD_CHANNEL }} \
          -m "Successfully built ${{ github.repository }}."
```
