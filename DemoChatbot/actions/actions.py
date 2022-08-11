# This files contains your custom actions which can be used to run
# custom Python code.
#
# See this guide on how to implement these action:
# https://rasa.com/docs/rasa/custom-actions


# This is a simple example for a custom action which utters "Hello World!"
# from mysql_connect import DataUpdate
from typing import Any, Text, Dict, List
import datetime as dt
from rasa_sdk import Action, Tracker
from rasa_sdk.executor import CollectingDispatcher
from rasa_core.events import UserUtteranceReverted

class ActionHelloWorld(Action):

    def name(self) -> Text:
        return "action_time"

    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker,
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:
        txt=tracker.latest_message.text
        print("txt",txt)
        dispatcher.utter_message(text=f"")
        return []


class ActionFallback(Action):

    def name(self) -> Text:
        return "fallback"

    def run(self, dispatcher: CollectingDispatcher,
            tracker: Tracker,
            domain: Dict[Text, Any]) -> List[Dict[Text, Any]]:

        dispatcher.utter_message(text=f"Sorrry didn't get that , try again")
        return []
