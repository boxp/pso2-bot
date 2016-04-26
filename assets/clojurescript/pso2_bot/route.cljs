(ns pso2-bot.route
  (:require [secretary.core :as secretary :refer-macros [defroute]]
            [om.core :as om :include-macros true]
            [om.dom :as dom :include-macros true]
            [goog.events :as events]
            [goog.history.EventType :as EventType]
            [pso2-bot.components.index :refer [index-page-view]]
            [pso2-bot.components.mypage :refer [my-page-view]])
  (:import goog.History))

(enable-console-print!)

(def application 
  (. js/document (getElementById "app")))

(secretary/set-config! :prefix "#")

(defroute index "/" []
  (om/root index-page-view {}
           {:target application}))

(defroute mypage "/mypage" []
  (om/root my-page-view {}
           {:target application}))

(defroute login "/login" [])

(defroute register "/register" [])

(let [h (History.)]
  (goog.events/listen h EventType/NAVIGATE #(secretary/dispatch! (.-token %)))
  (doto h (.setEnabled true)))
