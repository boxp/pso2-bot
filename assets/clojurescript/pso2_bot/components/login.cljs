(ns pso2-bot.components.signin
  (:require [sablono.core :as html :refer-macros [html]]
            [om-tools.core :refer-macros [defcomponent]]
            [om-tools.dom :include-macros true]
            [om.core :as om :include-macros true]))
            
(defcomponent signin-page-view [_ owner]
  (render [_]
    (html [:div.container
            [:form.form-signin
              [:h2 "Twitterアカウントで認証してください。"]]])))
