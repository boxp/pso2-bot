(ns pso2-bot.components.index
  (:require [sablono.core :as html :refer-macros [html]]
            [om-tools.core :refer-macros [defcomponent]]
            [om-tools.dom :include-macros true]
            [om.core :as om :include-macros true]
            [pso2-bot.components.common :refer [navigation-view]]))

(defcomponent index-page-view [_ owner]
  (render [_]
    (html [:div
           (om/build navigation-view {:active "Top"})
           [:div.container
             [:div.starter-tempate
               [:h1 "index"]]]])))
