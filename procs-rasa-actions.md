# procs-rasa-actions.md
⊕ [Actions](https://www.rasa.com/docs/core/customactions/)
⊕ [Events](https://www.rasa.com/docs/core/api/events/#events)

This request contains the next action as well as a lot of information about the conversation:

    next_action name of the predicted action that should be run
    sender_id   id of the conversation
    tracker serialised state of the conversations tracker
    domain  configuration of the domain

## Action Response Format
As a response to the action call from Core, you can modify the tracker, e.g. by setting slots and send responses back to the user. All of the modifications are done using events.

Here is an example json response:

```json
{
  "events": [
    {
      "event": "slot",
      "timestamp": null,
      "name": "concerts",
      "value": [
        {"artist": "Foo Fighters", "reviews": 4.5},
        {"artist": "Katy Perry", "reviews": 5.0}
      ]
    }
  ],
  "responses": [
    {"text": "Foo Fighters, Katy Perry"}
  ]
}
```
