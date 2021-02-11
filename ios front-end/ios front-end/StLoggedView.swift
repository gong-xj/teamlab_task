
import SwiftUI

struct StLoggedView: View {
    @State var id: String
    @State var name: String
    @State var scData: [Sc]
    @State var stOrTe: String
    
    var body: some View {
        TabView {
            ScView(scData:scData, name:name)
                .tabItem {
                     Image(systemName: "list.bullet")
                     Text("成績")
                }
            
            MeView(name:name, id:id, stOrTe:stOrTe)
                .tabItem {
                    Image(systemName: "person.fill")
                    Text("個人情報")
                }
        }
    }
}

struct StLoggedView_Previews: PreviewProvider {
    static var previews: some View {
        StLoggedView(id:"id??", name:"name??", scData:[], stOrTe:"学生??")
    }
}
