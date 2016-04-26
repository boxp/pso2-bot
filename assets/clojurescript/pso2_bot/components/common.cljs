(ns pso2-bot.components.common
  (:require [sablono.core :as html :refer-macros [html]]
            [om-tools.core :refer-macros [defcomponent]]
            [om-tools.dom :include-macros true]
            [om.core :as om :include-macros true]))

(defn- active?
  [me data]
  (if (= (:active data) me) 
    "active"))

(defcomponent navigation-view [data _]
  (render [_]
    (html [:nav {:class "navbar navbar-inverse navbar-fixed-top"}
            [:div.container
              [:div#navbar-header
                [:button {:type "button"
                          :class "navbar-toggle collapsed"
                          :data-toggle "collapse"
                          :data-target "#navbar"
                          :aria-expanded "false"
                          :aria-controls "navbar"}
                  [:span.sr-only "Toggle navigation"]
                  [:span.icon-bar]
                  [:span.icon-bar]
                  [:span.icon-bar]]
                [:a {:class "navbar-brand"
                     :href "#/"}
                    "pso2-bot"]]
              [:div#navbar {:class "collapse navbar-collapse"}
                [:ul {:class "nav navbar-nav"}
                  [:li {:class (active? "Top" data)}
                    [:a {:href "#/"} "Top"]]
                  [:li {:class (active? "MyPage" data)}
                    [:a {:href "#/mypage"} "MyPage"]]]]]])))
